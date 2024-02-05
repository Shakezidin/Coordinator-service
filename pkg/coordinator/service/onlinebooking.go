package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/Shakezidin/pkg/coordinator/client/pb"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
	msg "github.com/Shakezidin/pkg/rabbitmq"
	"github.com/google/uuid"
	"github.com/razorpay/razorpay-go"
)

// OnlinePaymentSVC handles online payments for bookings.
func (c *CoordinatorSVC) OnlinePaymentSVC(ctx context.Context, p *cpb.Booking) (*cpb.OnlinePaymentResponse, error) {
	// Retrieve user ID from Redis
	userIDKey := fmt.Sprintf("userId:%s", p.RefId)
	userId, err := c.redis.Get(ctx, userIDKey).Int()
	if err != nil {
		log.Printf("error getting user ID from Redis: %v", err)
		return nil, err
	}

	// Initialize Razorpay client
	client := razorpay.NewClient(c.cfg.RAZORPAYKEYID, c.cfg.RAZORPAYSECRETKEY)

	// Retrieve total fare from Redis
	amountKey := fmt.Sprintf("amount:%s", p.RefId)
	totalFare, err := c.redis.Get(ctx, amountKey).Int64()
	if err != nil {
		log.Printf("error getting total fare from Redis: %v", err)
		return nil, err
	}

	// Check if total fare is less than 1
	if totalFare < 1 {
		log.Println("total fare must be greater than or equal to 1")
		return nil, fmt.Errorf("total fare must be greater than or equal to 1")
	}

	// Calculate the amount based on payment type
	var amount float64
	if p.Typ == "full" {
		amount = float64(totalFare) * 100 // Convert total fare to paise
	} else if p.Typ == "advance" {
		amount = float64(totalFare) * 100 * 0.3 // Pay 30% advance
	}

	// Create order on Razorpay
	orderData := map[string]interface{}{
		"amount":   amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	order, err := client.Order.Create(orderData, nil)
	if err != nil {
		log.Printf("error creating order on Razorpay: %v", err)
		return nil, err
	}

	// Extract order ID from the response
	orderID, ok := order["id"].(string)
	if !ok {
		log.Println("error extracting order ID from Razorpay response")
		return nil, fmt.Errorf("failed to extract order ID from Razorpay response")
	}

	// Construct the response object
	response := &cpb.OnlinePaymentResponse{
		UserId:           int32(userId),
		TotalFare:        float32(amount) / 100,
		BookingReference: p.RefId,
		OrderId:          orderID,
	}

	return response, nil
}

func (c *CoordinatorSVC) PaymentConfirmedSVC(ctx context.Context, p *cpb.PaymentConfirmedRequest) (*cpb.BookingResponce, error) {
	// Start a new database transaction
	db := c.Repo.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			// If there's a panic, rollback the transaction
			tx.Rollback()
			panic(r)
		} else if err := recover(); err != nil {
			// If there's an error, rollback the transaction
			tx.Rollback()
			panic(err)
		}
	}()
	email_key := fmt.Sprintf("email:%s", p.ReferenceID)
	username_key := fmt.Sprintf("username:%s", p.ReferenceID)
	traveller_key := fmt.Sprintf("traveller:%s", p.ReferenceID)
	activity_key := fmt.Sprintf("activity_bookings:%s", p.ReferenceID)
	amount_key := fmt.Sprintf("amount:%s", p.ReferenceID)
	pkg_key := fmt.Sprintf("package:%v", p.ReferenceID)
	userIDKey := fmt.Sprintf("userId:%s", p.ReferenceID)
	userId, err := c.redis.Get(ctx, userIDKey).Int()
	if err != nil {
		log.Printf("error getting user ID from Redis: %v", err)
		return nil, err
	}

	email := c.redis.Get(ctx, email_key).Val()
	name := c.redis.Get(ctx, username_key).Val()
	total, _ := strconv.Atoi(p.Total)
	amoundata, err := c.redis.Get(ctx, amount_key).Int()
	rPay := dom.RazorPay{
		UserID:          uint(userId),
		RazorPaymentID:  p.PaymentId,
		Signature:       p.Signature,
		RazorPayOrderID: p.OrderID,
		AmountPaid:      float64(total),
	}

	var pkg dom.Package
	pkgData := c.redis.Get(ctx, pkg_key).Val()
	err = json.Unmarshal([]byte(pkgData), &pkg)
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error unmarshaling json err: %v", err.Error())
	}

	var activityBooking []dom.ActivityBooking
	activityData := c.redis.Get(ctx, activity_key).Val()

	err = json.Unmarshal([]byte(activityData), &activityBooking)
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error unmarshaling json err: %v", err.Error())
	}

	var travellers []dom.Traveller
	travellerData := c.redis.Get(ctx, traveller_key).Val()
	err = json.Unmarshal([]byte(travellerData), &travellers)
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error unmarshaling json err: %v", err.Error())
	}

	// Update the package within the transaction
	err = tx.Model(&dom.Package{}).Where("id = ?", pkg.ID).Updates(&pkg).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error updating package: %v", err.Error())
	}

	// Create travellers within the transaction
	for _, traveller := range travellers {
		err := tx.Create(&traveller).Error
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponce{
				Status: "fail",
			}, fmt.Errorf("error creating traveller: %v", err.Error())
		}
	}

	err = tx.Create(&rPay).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error creating razorpay: %v", err.Error())
	}

	// Create activity bookings within the transaction
	for _, activity := range activityBooking {
		err := tx.Create(&activity).Error
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponce{
				Status: "fail",
			}, fmt.Errorf("error creating activity: %v", err.Error())
		}
	}

	// Create the booking within the transaction
	var bookingID = uuid.New().String()
	var booking dom.Booking
	booking.BookingStatus = "success"
	booking.Bookings = travellers
	fmt.Println(total, "hhhhhh", float64(amoundata)*0.3)
	if float64(total) <= (float64(amoundata)*0.3)+5 {
		booking.PaymentMode = "advance"
	} else {
		booking.PaymentMode = "full amount"
	}
	booking.TotalPrice = amoundata
	booking.UserId = uint(userId)
	booking.BookingId = bookingID
	booking.Bookings = travellers
	booking.PackageId = pkg.ID
	booking.BookDate = time.Now()
	booking.StartDate = pkg.StartDate

	err = tx.Create(&booking).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error creating booking: %v", err.Error())
	}
	if float64(total) != float64(amoundata)*0.3 {
		coordinator := c.FindCoordinatorByPackageId(pkg.ID)

		reslt := tx.Model(&coordinator).Update("wallet", coordinator.Wallet+float64(amoundata)*0.70)
		if reslt.Error != nil {
			tx.Rollback()
			return &cpb.BookingResponce{
				Status: "fail",
			}, fmt.Errorf("error creating booking: %v", err.Error())
		}

		_, err := c.client.AdminAddWalletRequest(ctx, &pb.AdminAddWallet{
			Amount: float32(amoundata) * 0.30,
		})
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponce{
				Status: "fail",
			}, fmt.Errorf("error creating booking: %v", err.Error())
		}
	} else {
		_, err := c.client.AdminAddWalletRequest(ctx, &pb.AdminAddWallet{
			Amount: float32(amoundata),
		})
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponce{
				Status: "fail",
			}, fmt.Errorf("error creating booking: %v", err.Error())
		}
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	var message msg.Messages
	message.Amount = amoundata
	message.Email = email
	message.Username = name
	message.Messages = "Your booking is confirmed. Amount: "
	message.Subject = "Booking confirmed"

	msg.PublishConfirmationMessage(message)

	return &cpb.BookingResponce{
		Status:     "true",
		Booking_Id: bookingID,
	}, nil
}

func (c *CoordinatorSVC) FindCoordinatorByPackageId(pkgId uint) dom.User {
	pkg, _ := c.Repo.FetchPackage(pkgId)

	coordinator, _ := c.Repo.FetchUserById(pkg.ID)
	return *coordinator
}

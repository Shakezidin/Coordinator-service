package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	userIDKey := fmt.Sprintf("userId:%s", p.Ref_ID)
	userId, err := c.redis.Get(ctx, userIDKey).Int()
	if err != nil {
		return nil, err
	}

	// Initialize Razorpay client
	client := razorpay.NewClient(c.cfg.RAZORPAYKEYID, c.cfg.RAZORPAYSECRETKEY)

	// Retrieve total fare from Redis
	amountKey := fmt.Sprintf("amount:%s", p.Ref_ID)
	totalFare, err := c.redis.Get(ctx, amountKey).Int64()
	if err != nil {
		return nil, err
	}

	// Check if total fare is less than 1
	if totalFare < 1 {
		return nil, errors.New("total fare must be greater than or equal to 1")
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
		return nil, err
	}

	// Extract order ID from the response
	orderID, ok := order["id"].(string)
	if !ok {
		return nil, errors.New("failed to extract order ID from Razorpay response")
	}

	// Construct the response object
	response := &cpb.OnlinePaymentResponse{
		User_ID:           int32(userId),
		Total_Fare:        float32(amount) / 100,
		Booking_Reference: p.Ref_ID,
		Order_ID:          orderID,
	}

	return response, nil
}

// PaymentConfirmedSVC confirms payment and processes the booking.
func (c *CoordinatorSVC) PaymentConfirmedSVC(ctx context.Context, p *cpb.PaymentConfirmedRequest) (*cpb.BookingResponse, error) {
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

	emailKey := fmt.Sprintf("email:%s", p.Reference_ID)
	usernameKey := fmt.Sprintf("username:%s", p.Reference_ID)
	travellerKey := fmt.Sprintf("traveller:%s", p.Reference_ID)
	activityKey := fmt.Sprintf("activity_bookings:%s", p.Reference_ID)
	amountKey := fmt.Sprintf("amount:%s", p.Reference_ID)
	pkgKey := fmt.Sprintf("package:%v", p.Reference_ID)
	userIDKey := fmt.Sprintf("userId:%s", p.Reference_ID)

	userId, err := c.redis.Get(ctx, userIDKey).Int()
	if err != nil {
		return nil, err
	}

	email := c.redis.Get(ctx, emailKey).Val()
	name := c.redis.Get(ctx, usernameKey).Val()
	total, _ := strconv.ParseFloat(p.Total, 64)

	amountData, err := c.redis.Get(ctx, amountKey).Int()
	if err != nil {
		return nil, err
	}

	rPay := dom.RazorPay{
		UserID:          uint(userId),
		RazorPaymentID:  p.Payment_ID,
		Signature:       p.Signature,
		RazorPayOrderID: p.Order_ID,
		AmountPaid:      total,
	}

	var pkg dom.Package
	pkgData := c.redis.Get(ctx, pkgKey).Val()
	err = json.Unmarshal([]byte(pkgData), &pkg)
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponse{Status: "fail"}, err
	}

	var activityBooking []dom.ActivityBooking
	activityData := c.redis.Get(ctx, activityKey).Val()
	err = json.Unmarshal([]byte(activityData), &activityBooking)
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponse{Status: "fail"}, err
	}

	var travellers []dom.Traveller
	travellerData := c.redis.Get(ctx, travellerKey).Val()
	err = json.Unmarshal([]byte(travellerData), &travellers)
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponse{Status: "fail"}, err
	}

	// Update the package within the transaction
	err = tx.Model(&dom.Package{}).Where("id = ?", pkg.ID).Updates(&pkg).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponse{Status: "fail"}, err
	}

	// Create travellers within the transaction
	for _, traveller := range travellers {
		err := tx.Create(&traveller).Error
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponse{Status: "fail"}, err
		}
	}

	err = tx.Create(&rPay).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponse{Status: "fail"}, err
	}

	// Create activity bookings within the transaction
	for _, activity := range activityBooking {
		err := tx.Create(&activity).Error
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponse{Status: "fail"}, err
		}
	}

	// Create the booking within the transaction
	var bookingID = uuid.New().String()
	var booking dom.Booking
	booking.BookingStatus = "success"
	booking.Bookings = travellers
	if amountData != int(total) {
		booking.PaymentMode = "advance"
	} else {
		booking.PaymentMode = "full amount"
	}

	var codID uint
	if float64(total) != float64(amountData)*0.3 {
		coordinator := c.FindCoordinatorByPackageId(pkg.ID)
		codID = coordinator.ID

		result := tx.Model(&coordinator).Update("wallet", coordinator.Wallet+float64(amountData)*0.70)
		if result.Error != nil {
			tx.Rollback()
			return &cpb.BookingResponse{Status: "fail"}, err
		}
		_, err := c.client.AdminAddWalletRequest(ctx, &pb.AdminAddWallet{
			Amount: float32(amountData) * 0.3,
		})
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponse{Status: "fail"}, err
		}
	} else {
		coordinator := c.FindCoordinatorByPackageId(pkg.ID)
		codID = coordinator.ID

		_, err := c.client.AdminAddWalletRequest(ctx, &pb.AdminAddWallet{
			Amount: float32(amountData) * 0.3,
		})
		if err != nil {
			tx.Rollback()
			return &cpb.BookingResponse{Status: "fail"}, err
		}
	}

	booking.PaidPrice = int(total)
	booking.PackagePrice = amountData
	booking.UserID = uint(userId)
	booking.BookingID = bookingID
	booking.Bookings = travellers
	booking.PackageID = pkg.ID
	booking.BookDate = time.Now()
	booking.StartDate = pkg.StartDate
	booking.CoordinatorID = codID
	booking.UserEmail = email
	booking.CategoryID = pkg.Category.ID

	err = tx.Create(&booking).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponse{Status: "fail"}, err
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	var message msg.Messages
	message.Amount = amountData
	message.Email = email
	message.Username = name
	message.Messages = "Your booking is confirmed. Amount: "
	message.Subject = "Booking confirmed"

	go msg.PublishConfirmationMessage(message)

	return &cpb.BookingResponse{
		Status:     "true",
		Booking_ID: bookingID,
	}, nil
}

// FindCoordinatorByPackageId finds the coordinator by package ID.
func (c *CoordinatorSVC) FindCoordinatorByPackageId(pkgID uint) dom.User {
	pkg, _ := c.Repo.FetchPackage(pkgID)

	coordinator, _ := c.Repo.FetchUserById(pkg.ID)
	return *coordinator
}

package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
	"github.com/google/uuid"
)

const (
	redisKeyPrefix   = "traveller_details:"
	redisExpiration  = 10 * time.Minute
	redisRetryPolicy = 3
)

// TravellerDetails processes traveller details and stores them in Redis.
func (c *CoordinatorSVC) TravellerDetails(p *cpb.TravellerRequest) (*cpb.TravellerResponse, error) {
	ctx := context.Background()

	// Fetch package from repository
	pkgId, _ := strconv.Atoi(p.PackageId)
	pkg, err := c.Repo.FetchPackage(uint(pkgId))
	if err != nil {
		log.Print("package not found")
		return nil, errors.New("package not found")
	}

	if pkg.Availablespace < len(p.TravellerDetails) {
		log.Print("package space is not enough")
		return nil, errors.New("package have no space")
	} else {
		pkg.Availablespace = pkg.Availablespace - len(p.TravellerDetails)
	}

	var travellers []dom.Traveller
	var activityBookings []dom.ActivityBooking

	packageID, _ := strconv.Atoi(p.PackageId)
	userID, _ := strconv.Atoi(p.UserId)

	for _, travellerDetail := range p.TravellerDetails {
		var traveller dom.Traveller
		traveller.ID = uint(uuid.New().ID())
		traveller.Age = travellerDetail.Age
		traveller.Gender = travellerDetail.Gender
		traveller.Name = travellerDetail.Name
		traveller.PackageId = uint(packageID)
		traveller.UserId = uint(userID)

		travellers = append(travellers, traveller)

		for _, activityID := range travellerDetail.ActivityId {
			var activityBooking dom.ActivityBooking
			actId, _ := strconv.Atoi(activityID)
			activityBooking.ActivityId = uint(actId)
			activityBooking.TravellerId = traveller.ID

			activityBookings = append(activityBookings, activityBooking)
		}
	}
	activityTotal := c.calculateActivityTotal(p.TravellerDetails)
	refId := generateBookingReference()
	traveller_key := fmt.Sprintf("traveller%d", refId)
	activity_key := fmt.Sprintf("activity_bookings%d", refId)
	amount_key := fmt.Sprintf("amount%d", refId)
	pkg_key := fmt.Sprintf("package%d", refId)
	UserId_Key := fmt.Sprintf("userId%d", refId)

	err = c.storeInRedis(ctx, UserId_Key, p.UserId)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}

	err = c.storeInRedis(ctx, pkg_key, pkg)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}

	err = c.storeInRedis(ctx, traveller_key, travellers)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}

	err = c.storeInRedis(ctx, activity_key, activityBookings)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}

	err = c.storeInRedis(ctx, amount_key, pkg.MinPrice+activityTotal)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}

	return &cpb.TravellerResponse{
		PackagePrice:       int64(pkg.MinPrice),
		ActivityTotalPrice: int64(activityTotal),
		TotalPrice:         int64(pkg.MinPrice + activityTotal),
		RefId:              refId,
	}, nil
}

// Helper function to store data in Redis
func (c *CoordinatorSVC) storeInRedis(ctx context.Context, key string, data interface{}) error {
	marshaledData, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshaling %s: %v", key, err)
		return err
	}

	err = c.redis.Set(ctx, key, marshaledData, redisExpiration).Err()
	if err != nil {
		log.Printf("error storing %s in Redis: %v", key, err)
		return err
	}

	return nil
}

// Helper function to calculate total activity price
func (c *CoordinatorSVC) calculateActivityTotal(travellerDetails []*cpb.TravellerDetails) int {
	var activityTotal int

	for _, details := range travellerDetails {
		for _, activityID := range details.ActivityId {
			activityIDInt, _ := strconv.Atoi(activityID)
			activity, _ := c.Repo.FecthActivity(uint(activityIDInt))
			activityTotal += activity.Amount
		}
	}

	return activityTotal
}

func generateBookingReference() string {
	ref := uuid.New()
	return ref.String()
}

func (c *CoordinatorSVC) OfflineBooking(ctx context.Context, p *cpb.Booking) (*cpb.BookingResponce, error) {
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

	traveller_key := fmt.Sprintf("traveller%d", p.RefId)
	activity_key := fmt.Sprintf("activity_bookings%d", p.RefId)
	amount_key := fmt.Sprintf("amount%d", p.RefId)
	pkg_key := fmt.Sprintf("package%d", p.RefId)

	var pkg dom.Package
	pkgData := c.redis.Get(ctx, pkg_key).Val()
	err := json.Unmarshal([]byte(pkgData), &pkg)
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

	amoundata, err := c.redis.Get(ctx, amount_key).Int()

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
	booking.PaymentMode = "offline"
	booking.TotalPrice = amoundata
	booking.UserId = uint(p.UserId)
	booking.BookingId = bookingID
	booking.Bookings = travellers
	booking.PackageId = pkg.ID
	booking.Activities = activityBooking

	err = tx.Create(&booking).Error
	if err != nil {
		tx.Rollback()
		return &cpb.BookingResponce{
			Status: "fail",
		}, fmt.Errorf("error creating booking: %v", err.Error())
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	return &cpb.BookingResponce{
		Booking_Id: bookingID,
	}, nil
}

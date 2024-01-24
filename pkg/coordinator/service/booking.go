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
		return nil, err
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

	err = c.storeInRedis(ctx, pkg_key, pkg)
	if err != nil {
		return nil, err
	}

	err = c.storeInRedis(ctx, traveller_key, travellers)
	if err != nil {
		return nil, err
	}

	err = c.storeInRedis(ctx, activity_key, activityBookings)
	if err != nil {
		return nil, err
	}

	err = c.storeInRedis(ctx, amount_key, pkg.MinPrice+activityTotal)
	if err != nil {
		return nil, err
	}

	return &cpb.TravellerResponse{
		Status:             "success",
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

	traveller_key := fmt.Sprintf("traveller%d", p.RefId)
	activity_key := fmt.Sprintf("activity_bookings%d", p.RefId)
	amount_key := fmt.Sprintf("amount%d", p.RefId)
	pkg_key := fmt.Sprintf("package%d", p.RefId)

	var pkg dom.Package
	pkgData := c.redis.Get(ctx, pkg_key).Val()
	err := json.Unmarshal([]byte(pkgData), &pkg)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json err: %v", err.Error())
	}

	var activityBooking []dom.ActivityBooking
	activityData := c.redis.Get(ctx, activity_key).Val()

	err = json.Unmarshal([]byte(activityData), &activityBooking)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json err: %v", err.Error())
	}

	amoundata, err := c.redis.Get(ctx, amount_key).Int()

	var travellers []dom.Traveller
	travellerData := c.redis.Get(ctx, traveller_key).Val()
	err = json.Unmarshal([]byte(travellerData), &travellers)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json err: %v", err.Error())
	}

	err = c.Repo.UpdatePackage(&pkg)
	if err != nil {
		return nil, errors.New("error while package updating")
	}
	for _, traveller := range travellers {
		err := c.Repo.CreateTraveller(traveller)
		if err != nil {
			return nil, err
		}
	}

	for _, activity := range activityBooking {
		err := c.Repo.CreateActivityBooking(activity)
		if err != nil {
			return nil, err
		}
	}

	var bookingID = uuid.New().String()

	var booking dom.Booking
	booking.BookingStatus = "success"
	booking.Bookings = travellers
	booking.PaymentMode = "offline"
	booking.TotalPrice = amoundata
	booking.UserId = uint(p.UserId)
	booking.BookingId = bookingID

	err = c.Repo.CreateBooking(booking)
	if err != nil {
		return nil, err
	}
	return &cpb.BookingResponce{
		Status:     "Success",
		Booking_Id: bookingID,
	}, nil
}

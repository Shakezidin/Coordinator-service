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

	fmt.Println(pkg.AvailableSpace ,"hhhhh",len(p.TravellerDetails))
	if pkg.AvailableSpace < len(p.TravellerDetails) {
		log.Print("package space is not enough")
		return nil, errors.New("package have no space")
	} else {
		pkg.AvailableSpace = pkg.AvailableSpace - len(p.TravellerDetails)
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
	email_key := fmt.Sprintf("email:%s", refId)
	username_key := fmt.Sprintf("username:%s", refId)
	traveller_key := fmt.Sprintf("traveller:%s", refId)
	activity_key := fmt.Sprintf("activity_bookings:%s", refId)
	amount_key := fmt.Sprintf("amount:%s", refId)
	pkg_key := fmt.Sprint("package:", refId)
	UserId_Key := fmt.Sprintf("userId:%s", refId)

	err = c.storeInRedis(ctx, email_key, p.Email)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}
	err = c.storeInRedis(ctx, username_key, p.Username)
	if err != nil {
		return nil, errors.New("error while storing to redis")
	}
	userid, _ := strconv.Atoi(p.UserId)
	err = c.storeInRedis(ctx, UserId_Key, userid)
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

	totalAmount := pkg.MinPrice + activityTotal
	advanceAmount := float64(totalAmount) * 0.3

	return &cpb.TravellerResponse{
		PackagePrice:       int64(pkg.MinPrice),
		ActivityTotalPrice: int64(activityTotal),
		TotalPrice:         int64(totalAmount),
		AdvanceAmount:      int64(advanceAmount),
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

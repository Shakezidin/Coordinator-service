package service

import (
	"context"
	"encoding/json"
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

	// Marshal and store traveller details and activity bookings in Redis
	err = c.storeInRedis(ctx, "travellers", travellers)
	if err != nil {
		return nil, err
	}

	err = c.storeInRedis(ctx, "activity_bookings", activityBookings)
	if err != nil {
		return nil, err
	}

	// Calculate total activity price
	activityTotal := c.calculateActivityTotal(p.TravellerDetails)

	return &cpb.TravellerResponse{
		Status:             "success",
		PackagePrice:       int64(pkg.MinPrice),
		ActivityTotalPrice: int64(activityTotal),
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

package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	pkgID, err := strconv.Atoi(p.PackageId)
	if err != nil {
		return nil, errors.New("invalid package ID")
	}
	pkg, err := c.Repo.FetchPackage(uint(pkgID))
	if err != nil {
		return nil, errors.New("package not found")
	}

	// Check if package has enough space for travellers
	if pkg.AvailableSpace < len(p.TravellerDetails) {
		return nil, errors.New("package has insufficient space for travellers")
	}
	pkg.AvailableSpace -= len(p.TravellerDetails)

	// Prepare traveller and activity booking data
	var travellers []dom.Traveller
	var activityBookings []dom.ActivityBooking

	userID, _ := strconv.Atoi(p.UserId)

	for _, travellerDetail := range p.TravellerDetails {
		traveller := dom.Traveller{
			ID:        uint(uuid.New().ID()),
			Age:       travellerDetail.Age,
			Gender:    travellerDetail.Gender,
			Name:      travellerDetail.Name,
			PackageID: uint(pkgID),
			UserID:    uint(userID),
		}
		travellers = append(travellers, traveller)

		for _, activityID := range travellerDetail.ActivityId {
			actID, err := strconv.Atoi(activityID)
			if err != nil {
				continue
			}
			activityBooking := dom.ActivityBooking{
				ActivityID:  uint(actID),
				TravellerID: traveller.ID,
			}
			activityBookings = append(activityBookings, activityBooking)
		}
	}

	// Calculate total amount
	activityAmount := c.CalculateActivityTotal(p.TravellerDetails)
	advanceAmount := float64(pkg.MinPrice+activityAmount) * 0.3

	// Store data in Redis
	refID := generateBookingReference()
	err = c.StoreTravellerDetailsInRedis(ctx, refID, p, pkg, travellers, activityBookings, pkg.MinPrice+activityAmount)
	if err != nil {
		return nil, errors.New("error storing traveller details")
	}

	// Return response
	return &cpb.TravellerResponse{
		PackagePrice:       int64(pkg.MinPrice),
		ActivityTotalPrice: int64(activityAmount),
		TotalPrice:         int64(pkg.MinPrice + activityAmount),
		AdvanceAmount:      int64(advanceAmount),
		RefId:              refID,
	}, nil
}

// Helper function to store traveller details in Redis
func (c *CoordinatorSVC) StoreTravellerDetailsInRedis(ctx context.Context, refID string, p *cpb.TravellerRequest, pkg *dom.Package, travellers []dom.Traveller, activityBookings []dom.ActivityBooking, totalAmount int) error {
	emailKey := fmt.Sprintf("email:%s", refID)
	usernameKey := fmt.Sprintf("username:%s", refID)
	userIDKey := fmt.Sprintf("userId:%s", refID)
	packageKey := fmt.Sprintf("package:%s", refID)
	travellerKey := fmt.Sprintf("traveller:%s", refID)
	activityKey := fmt.Sprintf("activity_bookings:%s", refID)
	amountKey := fmt.Sprintf("amount:%s", refID)

	err := c.StoreInRedis(ctx, emailKey, p.Email)
	if err != nil {
		return fmt.Errorf("error storing email in Redis: %v", err)
	}
	err = c.StoreInRedis(ctx, usernameKey, p.Username)
	if err != nil {
		return fmt.Errorf("error storing username in Redis: %v", err)
	}
	userID, err := strconv.Atoi(p.UserId)
	if err != nil {
		return fmt.Errorf("invalid user ID: %s", p.UserId)
	}
	err = c.StoreInRedis(ctx, userIDKey, userID)
	if err != nil {
		return fmt.Errorf("error storing user ID in Redis: %v", err)
	}
	err = c.StoreInRedis(ctx, packageKey, pkg)
	if err != nil {
		return fmt.Errorf("error storing package details in Redis: %v", err)
	}
	err = c.StoreInRedis(ctx, travellerKey, travellers)
	if err != nil {
		return fmt.Errorf("error storing traveller details in Redis: %v", err)
	}
	err = c.StoreInRedis(ctx, activityKey, activityBookings)
	if err != nil {
		return fmt.Errorf("error storing activity details in Redis: %v", err)
	}
	err = c.StoreInRedis(ctx, amountKey, totalAmount)
	if err != nil {
		return fmt.Errorf("error storing amount details in Redis: %v", err)
	}

	return nil
}

// Helper function to store data in Redis
func (c *CoordinatorSVC) StoreInRedis(ctx context.Context, key string, data interface{}) error {
	marshaledData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = c.redis.Set(ctx, key, marshaledData, redisExpiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Helper function to calculate total activity price
func (c *CoordinatorSVC) CalculateActivityTotal(travellerDetails []*cpb.TravellerDetails) int {
	var activityTotal int

	for _, details := range travellerDetails {
		for _, activityID := range details.ActivityId {
			actID, err := strconv.Atoi(activityID)
			if err != nil {
				continue
			}
			activity, err := c.Repo.FetchActivity(uint(actID))
			if err != nil {
				continue
			}
			activityTotal += activity.Amount
		}
	}

	return activityTotal
}

func generateBookingReference() string {
	return uuid.New().String()
}

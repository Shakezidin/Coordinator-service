package service

import (
	"context"
	"errors"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/Shakezidin/pkg/entities/packages"
)

// ViewHistorySVC retrieves booking history.
func (c *CoordinatorSVC) ViewHistorySVC(p *cpb.View) (*cpb.Histories, error) {
	offset := 10 * (p.Page - 1)
	limit := 10
	var history []*packages.Booking
	var err error

	if p.Status == "false" {
		history, err = c.Repo.FetchHistory(int(offset), limit, uint(p.Id))
		if err != nil {
			return nil, errors.New("error while fetching package history")
		}
	} else {
		history, err = c.Repo.FetchBookings(int(offset), limit, uint(p.Id))
		if err != nil {
			return nil, errors.New("error while fetching booking history")
		}
	}

	var pbHistory []*cpb.History
	for _, h := range history {
		pbHistory = append(pbHistory, &cpb.History{
			Id:              int64(h.ID),
			PaymentMode:     h.PaymentMode,
			BookingStatus:   h.BookingStatus,
			CancelledStatus: h.CancelledStatus,
			TotalPrice:      int64(h.PackagePrice),
			UserId:          int64(h.UserID),
			BookingId:       h.BookingID,
			BookDate:        h.BookDate.Format("02-01-2006"),
			StartDate:       h.StartDate.Format("02-01-2006"),
			PaidAmount:      int64(h.PaidPrice),
		})
	}

	return &cpb.Histories{
		History: pbHistory,
	}, nil
}

// ViewBookingSVC retrieves a specific booking.
func (c *CoordinatorSVC) ViewBookingSVC(p *cpb.View) (*cpb.History, error) {
	booking, err := c.Repo.FetchBooking(uint(p.Id))
	if err != nil {
		return nil, errors.New("booking not found")
	}

	var travellers []*cpb.TravellerDetails
	for _, t := range booking.Bookings {
		travellers = append(travellers, &cpb.TravellerDetails{
			Id:     int64(t.ID),
			Name:   t.Name,
			Age:    t.Age,
			Gender: t.Gender,
		})
	}

	return &cpb.History{
		Id:              int64(booking.ID),
		PaymentMode:     booking.PaymentMode,
		BookingStatus:   booking.BookingStatus,
		CancelledStatus: booking.CancelledStatus,
		TotalPrice:      int64(booking.PackagePrice),
		UserId:          int64(booking.UserID),
		BookingId:       booking.BookingID,
		BookDate:        booking.BookDate.Format("02-01-2006"),
		StartDate:       booking.StartDate.Format("02-01-2006"),
		Travellers:      travellers,
		PaidAmount:      int64(booking.PaidPrice),
	}, nil
}

// CancelBookingSVC cancels a booking.
func (c *CoordinatorSVC) CancelBookingSVC(p *cpb.View) (*cpb.Response, error) {
	booking, err := c.Repo.FetchBooking(uint(p.Id))
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "booking not found",
		}, errors.New("booking not found")
	}

	if booking.CancelledStatus == "cancelled" {
		return &cpb.Response{
			Status:  "fail",
			Message: "booking already cancelled",
		}, errors.New("booking already cancelled")
	}

	booking.CancelledStatus = "cancelled"

	pkg, err := c.Repo.FetchPackage(booking.PackageID)
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "package not found",
		}, errors.New("package not found")
	}

	pkg.AvailableSpace += len(booking.Bookings)
	err = c.Repo.UpdatePackage(*pkg)
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "failed to update package",
		}, errors.New("failed to update package")
	}

	// Additional logic for coordinator wallet adjustment and admin notification goes here

	return &cpb.Response{
		Status:  "Success",
		Message: "booking cancelled successfully",
	}, nil
}

// ViewTravellerSVC retrieves details of a specific traveller.
func (c *CoordinatorSVC) ViewTravellerSVC(p *cpb.View) (*cpb.TravellerDetails, error) {
	traveller, err := c.Repo.FetchTraveller(uint(p.Id))
	if err != nil {
		return nil, errors.New("traveller not found")
	}

	activityBookings, _ := c.Repo.FetchActivityBookingofUser(uint(p.Id))

	var activities []*cpb.Activity
	for _, a := range activityBookings {
		activities = append(activities, &cpb.Activity{
			ActivityId:   int64(a.Activity.ID),
			Activityname: a.Activity.ActivityName,
			Description:  a.Activity.Description,
			Location:     a.Activity.Location,
			ActivityType: a.Activity.ActivityType,
			Amount:       int64(a.Activity.Amount),
			Date:         a.Activity.Date.Format("02-01-2006"),
			Time:         a.Activity.Time.Format("03:04 PM"),
		})
	}

	return &cpb.TravellerDetails{
		Name:     traveller.Name,
		Age:      traveller.Age,
		Gender:   traveller.Gender,
		Activity: activities,
	}, nil
}

// SearchBookingSVC searches for bookings based on criteria.
func (c *CoordinatorSVC) SearchBookingSVC(p *cpb.BookingSearchCriteria) (*cpb.Histories, error) {
	ctx := context.Background()

	bookings, err := c.Repo.SearchBookings(ctx, p)
	if err != nil {
		return nil, errors.New("error while searching bookings")
	}

	var histories []*cpb.History
	for _, booking := range bookings {
		history := &cpb.History{
			Id:              int64(booking.ID),
			PaymentMode:     booking.PaymentMode,
			BookingStatus:   booking.BookingStatus,
			CancelledStatus: booking.CancelledStatus,
			TotalPrice:      int64(booking.PackagePrice),
			UserId:          int64(booking.UserID),
			BookingId:       booking.BookingID,
			BookDate:        booking.BookDate.Format("02-01-2006"),
			StartDate:       booking.StartDate.Format("02-01-2006"),
			PaidAmount:      int64(booking.PaidPrice),
		}
		histories = append(histories, history)
	}

	return &cpb.Histories{
		History: histories,
	}, nil
}

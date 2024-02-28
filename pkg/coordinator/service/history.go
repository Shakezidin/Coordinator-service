package service

import (
	"context"
	"errors"

	pb "github.com/Shakezidin/pkg/coordinator/client/pb"
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
		history, err = c.Repo.FetchHistory(int(offset), limit, uint(p.ID))
		if err != nil {
			return nil, errors.New("error while fetching package history")
		}
	} else {
		history, err = c.Repo.FetchBookings(int(offset), limit, uint(p.ID))
		if err != nil {
			return nil, errors.New("error while fetching booking history")
		}
	}

	var pbHistory []*cpb.History
	for _, h := range history {
		pbHistory = append(pbHistory, &cpb.History{
			ID:               int64(h.ID),
			Payment_Mode:     h.PaymentMode,
			Booking_Status:   h.BookingStatus,
			Cancelled_Status: h.CancelledStatus,
			Total_Price:      int64(h.PackagePrice),
			User_ID:          int64(h.UserID),
			Booking_ID:       h.BookingID,
			Book_Date:        h.BookDate.Format("02-01-2006"),
			Start_Date:       h.StartDate.Format("02-01-2006"),
			Paid_Amount:      int64(h.PaidPrice),
		})
	}

	return &cpb.Histories{
		Histories: pbHistory,
	}, nil
}

// ViewBookingSVC retrieves a specific booking.
func (c *CoordinatorSVC) ViewBookingSVC(p *cpb.View) (*cpb.History, error) {
	booking, err := c.Repo.FetchBooking(uint(p.ID))
	if err != nil {
		return nil, errors.New("booking not found")
	}

	var travellers []*cpb.TravellerDetails
	for _, t := range booking.Bookings {
		travellers = append(travellers, &cpb.TravellerDetails{
			ID:     int64(t.ID),
			Name:   t.Name,
			Age:    t.Age,
			Gender: t.Gender,
		})
	}

	return &cpb.History{
		ID:               int64(booking.ID),
		Payment_Mode:     booking.PaymentMode,
		Booking_Status:   booking.BookingStatus,
		Cancelled_Status: booking.CancelledStatus,
		Total_Price:      int64(booking.PackagePrice),
		User_ID:          int64(booking.UserID),
		Booking_ID:       booking.BookingID,
		Book_Date:        booking.BookDate.Format("02-01-2006"),
		Start_Date:       booking.StartDate.Format("02-01-2006"),
		Travellers:       travellers,
		Paid_Amount:      int64(booking.PaidPrice),
	}, nil
}

// CancelBookingSVC cancels a booking.
func (c *CoordinatorSVC) CancelBookingSVC(p *cpb.View) (*cpb.Response, error) {
	booking, err := c.Repo.FetchBooking(uint(p.ID))
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
	coordinator, err := c.Repo.FetchUserById(pkg.CoordinatorID)
	if err != nil {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("error while fetching coordinator")
	}
	if booking.PaymentMode == "full amount" {
		coordinator.Wallet -= float64(booking.PackagePrice) * 0.70
		err = c.Repo.UpdateUser(coordinator)
		if err != nil {
			return &cpb.Response{
				Status: "fail",
			}, errors.New("error while updating coordinator")
		}
	}
	err = c.Repo.UpdateBooking(*booking)
	if err != nil {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("error while updating booking")
	}
	err = c.Repo.UpdatePackage(*pkg)
	if err != nil {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("error while updating package")
	}
	var ctx = context.Background()
	_, err = c.client.AdminReduseWalletRequesr(ctx, &pb.AdminAddWallet{
		Amount: float32(booking.PackagePrice) * 0.30,
	})

	if err != nil {
		return &cpb.Response{
			Status: "fail",
		}, err
	}

	return &cpb.Response{
		Amount:  int64(booking.PaidPrice),
		Status:  "Success",
		Message: "Package cancelled success",
	}, nil
}

// ViewTravellerSVC retrieves details of a specific traveller.
func (c *CoordinatorSVC) ViewTravellerSVC(p *cpb.View) (*cpb.TravellerDetails, error) {
	traveller, err := c.Repo.FetchTraveller(uint(p.ID))
	if err != nil {
		return nil, errors.New("traveller not found")
	}

	activityBookings, _ := c.Repo.FetchActivityBookingofUser(uint(p.ID))

	var activities []*cpb.Activity
	for _, a := range activityBookings {
		activities = append(activities, &cpb.Activity{
			Activity_ID:   int64(a.Activity.ID),
			Activity_Name: a.Activity.ActivityName,
			Description:   a.Activity.Description,
			Location:      a.Activity.Location,
			Activity_Type: a.Activity.ActivityType,
			Amount:        int64(a.Activity.Amount),
			Date:          a.Activity.Date.Format("02-01-2006"),
			Time:          a.Activity.Time.Format("03:04 PM"),
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
			ID:               int64(booking.ID),
			Payment_Mode:     booking.PaymentMode,
			Booking_Status:   booking.BookingStatus,
			Cancelled_Status: booking.CancelledStatus,
			Total_Price:      int64(booking.PackagePrice),
			User_ID:          int64(booking.UserID),
			Booking_ID:       booking.BookingID,
			Book_Date:        booking.BookDate.Format("02-01-2006"),
			Start_Date:       booking.StartDate.Format("02-01-2006"),
			Paid_Amount:      int64(booking.PaidPrice),
		}
		histories = append(histories, history)
	}

	return &cpb.Histories{
		Histories: histories,
	}, nil
}

package repository

import (
	"context"
	"time"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) CreateTraveller(traveller cDOM.Traveller) error {
	if err := c.DB.Create(&traveller).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) CreateActivityBooking(activity cDOM.ActivityBooking) error {
	if err := c.DB.Create(&activity).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) CreateBooking(booking cDOM.Booking) error {
	if err := c.DB.Create(&booking).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) UpdateBooking(booking cDOM.Booking) error {
	if err := c.DB.Save(&booking).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FetchNextDayTrip(date string) (*[]cDOM.Booking, error) {
	var booking []cDOM.Booking
	if err := c.DB.Preload("Package").Where("start_date = ?", date).Find(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *CoordinatorRepo) SearchBookings(ctx context.Context, criteria *cpb.BookingSearchCriteria) ([]*cDOM.Booking, error) {
	var bookings []*cDOM.Booking

	query := r.DB.Model(&cDOM.Booking{})

	// Apply filters based on the provided search criteria
	if criteria.PaymentMode != "" {
		query = query.Where("payment_mode = ?", criteria.PaymentMode)
	}
	if criteria.BookingStatus != "" {
		query = query.Where("booking_status = ?", criteria.BookingStatus)
	}
	if criteria.CancelledStatus {
		query = query.Where("cancelled_status = ?", true)
	}
	if criteria.UserEmail != "" {
		query = query.Where("user_email = ?", criteria.UserEmail)
	}
	if criteria.BookingId != "" {
		query = query.Where("booking_id = ?", criteria.BookingId)
	}
	if criteria.BookDate != "" {
		bookDate, err := time.Parse("02-01-2006", criteria.BookDate)
		if err != nil {
			return nil, err
		}
		query = query.Where("book_date >= ?", bookDate)
	}
	if criteria.StartDate != "" {
		startDate, err := time.Parse("02-01-2006", criteria.StartDate)
		if err != nil {
			return nil, err
		}
		query = query.Where("start_date >= ?", startDate)
	}
	if criteria.CoordinatorId != 0 {
		query = query.Where("coordinator_id = ?", criteria.CoordinatorId)
	}
	
	if criteria.CatageryId != 0 {
		query = query.Where("category_id = ?", criteria.CatageryId)
	}

	if err := query.Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

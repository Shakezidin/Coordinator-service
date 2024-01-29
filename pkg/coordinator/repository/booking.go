package repository

import (
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

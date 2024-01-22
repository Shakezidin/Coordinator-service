package repository

import (
	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) CreateTraveller(traveller cDOM.Traveller) error {
	if err := c.db.Create(&traveller).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) CreateActivityBooking(activity cDOM.ActivityBooking) error {
	if err := c.db.Create(&activity).Error; err != nil {
		return err
	}
	return nil
}

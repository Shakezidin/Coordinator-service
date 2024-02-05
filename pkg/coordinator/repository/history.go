package repository

import cDOM "github.com/Shakezidin/pkg/entities/packages"

func (c *CoordinatorRepo) FetchHistory(offset, limit int, id uint) (*[]cDOM.Booking, error) {
	var booking *[]cDOM.Booking
	if err := c.DB.Where("user_id = ?", id).Offset(offset).Limit(limit).Find(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (c *CoordinatorRepo) FetchBooking(id uint) (*cDOM.Booking, error) {
    booking := &cDOM.Booking{}
    if err := c.DB.Where("id = ?", id).Preload("Bookings").First(booking).Error; err != nil {
        return nil, err
    }
    return booking, nil
}


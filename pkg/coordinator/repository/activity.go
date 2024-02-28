package repository

import cDOM "github.com/Shakezidin/pkg/entities/packages"

func (c *CoordinatorRepo) CreateActivity(actvt *cDOM.Activity) error {
	if err := c.DB.Create(&actvt).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FecthDestinationActivity(id uint) ([]*cDOM.Activity, error) {
	var actvty []*cDOM.Activity
	if err := c.DB.Where("destination_id = ?", id).Find(&actvty).Error; err != nil {
		return nil, err
	}
	return actvty, nil
}

func (c *CoordinatorRepo) FetchActivityBookingofUser(id uint) ([]*cDOM.ActivityBooking, error) {
    var actvtyBooking []*cDOM.ActivityBooking
    if err := c.DB.Preload("Activity").Where("traveller_id = ?", id).Find(&actvtyBooking).Error; err != nil {
        return nil, err
    }
    return actvtyBooking, nil
}


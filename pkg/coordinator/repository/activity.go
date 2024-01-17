package repository

import cDOM "github.com/Shakezidin/pkg/entities/packages"

func (c *CoordinatorRepo) CreateActivity(actvt *cDOM.Activity) error {
	if err := c.db.Create(&actvt).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FecthDestinationActivity(id uint) ([]*cDOM.Activity, error) {
	var actvty []*cDOM.Activity
	if err := c.db.Where("destination_id = ?", id).Find(&actvty).Error; err != nil {
		return nil, err
	}
	return actvty, nil
}

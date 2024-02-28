package repository

import (
	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) CreateDestination(dtnt *cDOM.Destination) error {
	if err := c.DB.Create(&dtnt).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FetchPackageDestination(id uint) ([]*cDOM.Destination, error) {
	var dstn []*cDOM.Destination
	if err := c.DB.Where("package_id = ?", id).Find(&dstn).Error; err != nil {
		return nil, err
	}
	return dstn, nil
}

func (c *CoordinatorRepo) FecthDestination(id uint) (*cDOM.Destination, error) {
	var dstn cDOM.Destination
	if err := c.DB.Where("id = ?", id).First(&dstn).Error; err != nil {
		return nil, err
	}
	return &dstn, nil
}

func (c *CoordinatorRepo) FetchActivity(id uint) (*cDOM.Activity, error) {
	var activity *cDOM.Activity
	if err := c.DB.Where("id = ?", id).First(&activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}
package repository

import (
	"time"

	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) FindUnboundedPackages(offset, limit int, PickupPlace, Finaldestination string, MaxDestination int64, startDate, endDate time.Time) ([]*cDOM.Package, error) {
	var packages []*cDOM.Package
	if endDate.IsZero() {
		if err := c.DB.Offset(offset).Limit(limit).Where("trip_status = ? AND start_location = ? AND destination = ? AND num_of_destination <= ?", true,
			PickupPlace, Finaldestination, MaxDestination).Find(&packages).Error; err != nil {
			return nil, err
		}
		return packages, nil
	}

	if err := c.DB.Offset(offset).Limit(limit).Where("trip_status = ? AND start_location = ? AND destination = ? AND start_date = ? AND end_date = ? AND num_of_destination <= ?", true,
		PickupPlace, Finaldestination, startDate, endDate, MaxDestination).Find(&packages).Error; err != nil {
		return nil, err
	}
	return packages, nil
}

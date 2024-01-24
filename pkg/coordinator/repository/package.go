package repository

import (
	cDOM "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorRepo) FetchPackage(id uint) (*cDOM.Package, error) {
	var pkg cDOM.Package
	if err := c.db.Preload("Category").Where("id = ?", id).First(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (c *CoordinatorRepo) CreatePackage(pkg *cDOM.Package) error {
	return c.db.Create(&pkg).Error
}

func (c *CoordinatorRepo) FetchPackages(offset, limit int, val string) ([]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Offset(offset).Limit(limit).Where("trip_status = ?", val).Find(&packages).Error; err != nil {
		return nil, err
	}
	return packages, nil
}

func (c *CoordinatorRepo) FindCoordinatorPackages(id uint) (*[]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Where("coordinator_id = ?", id).Find(&packages).Error; err != nil {
		return nil, err
	}
	return &packages, nil
}

func (c *CoordinatorRepo) CreateCatagory(catagory cDOM.Category) error {
	if err := c.db.Create(&catagory).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FetchAllPackages(offset, limit int) (*[]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Offset(offset).Limit(limit).Find(&packages).Error; err != nil {
		return nil, err
	}
	return &packages, nil
}

func (c *CoordinatorRepo) PackageStatusUpdate(id uint) error {
	var packageToUpdate cDOM.Package
	if err := c.db.First(&packageToUpdate, id).Error; err != nil {
		return err
	}

	// Flip the status
	packageToUpdate.TripStatus = !packageToUpdate.TripStatus

	if err := c.db.Save(&packageToUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (c *CoordinatorRepo) FetchCatagories(offset, limit int) ([]*cDOM.Category, error) {
	var categories []*cDOM.Category
	if err := c.db.Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CoordinatorRepo) FetchCatagory(catagory string) (*cDOM.Category, error) {
	var catagories *cDOM.Category
	if err := c.db.First(&catagories).Error; err != nil {
		return nil, err
	}
	return catagories, nil
}

func (c *CoordinatorRepo) UpdatePackage(pkg *cDOM.Package) error {
	if err := c.db.Save(&pkg).Error; err != nil {
		return err
	}
	return nil
}

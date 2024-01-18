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
	if err := c.db.Create(&pkg).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FetchPackages(val string) (*[]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Where("trip_status = ?", val).Find(&packages).Error; err != nil {
		return nil, err
	}
	return &packages, nil
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

func (c *CoordinatorRepo) FetchAllPackages() (*[]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Find(&packages).Error; err != nil {
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

func (c CoordinatorRepo) FetchCatagories() ([]*cDOM.Category, error) {
	var catagories []*cDOM.Category
	if err := c.db.Find(&catagories).Error; err != nil {
		return nil, err
	}
	return catagories, nil
}

package repository

import (
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	cDOM "github.com/Shakezidin/pkg/entities/packages"
	"gorm.io/gorm"
)

type CoordinatorRepo struct {
	db *gorm.DB
}

func (c *CoordinatorRepo) SignupRepo(user *cDOM.User) error {
	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FindUserByEmail(email string) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) FindUserByPhone(number int) (*cDOM.User, error) {
	var user cDOM.User
	if err := c.db.Where("phone = ?", number).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CoordinatorRepo) CreateUser(user *cDOM.User) error {
	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FindCoordinatorPackages(id uint) (*[]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Where("coordinator_id = ?", id).Find(&packages).Error; err != nil {
		return nil, err
	}
	return &packages, nil
}

func (c *CoordinatorRepo) FetchAllPackages() (*[]cDOM.Package, error) {
	var packages []cDOM.Package
	if err := c.db.Where("trip_status = ?", true).Find(&packages).Error; err != nil {
		return nil, err
	}
	return &packages, nil
}

func (c *CoordinatorRepo) CreatePackage(pkg *cDOM.Package) error {
	if err := c.db.Create(&pkg).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) CreateDestination(dtnt *cDOM.Destination) error {
	if err := c.db.Create(&dtnt).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) CreateActivity(actvt *cDOM.Activity) error {
	if err := c.db.Create(&actvt).Error; err != nil {
		return err
	}
	return nil
}

func (c *CoordinatorRepo) FetchPackage(id uint) (*cDOM.Package, error) {
	var pkg cDOM.Package
	if err := c.db.Preload("Category").Where("id = ?", id).First(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (c *CoordinatorRepo) FetchPackageDestination(id uint) ([]*cDOM.Destination, error) {
	var dstn []*cDOM.Destination
	if err := c.db.Where("package_id = ?", id).Find(&dstn).Error; err != nil {
		return nil, err
	}
	return dstn, nil
}

func (c *CoordinatorRepo) FecthDestination(id uint) (*cDOM.Destination, error) {
	var dstn cDOM.Destination
	if err := c.db.Where("id = ?", id).First(&dstn).Error; err != nil {
		return nil, err
	}
	return &dstn, nil
}

func (c *CoordinatorRepo) FecthDestinationActivity(id uint) ([]*cDOM.Activity, error) {
	var actvty []*cDOM.Activity
	if err := c.db.Where("destination_id = ?", id).Find(&actvty).Error; err != nil {
		return nil, err
	}
	return actvty, nil
}

func (c *CoordinatorRepo) FecthActivity(id uint) (*cDOM.Activity, error) {
	var activity *cDOM.Activity
	if err := c.db.Where("id = ?", id).First(&activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

func NewCoordinatorRepo(db *gorm.DB) inter.CoordinatorRepoInter {
	return &CoordinatorRepo{
		db: db,
	}
}

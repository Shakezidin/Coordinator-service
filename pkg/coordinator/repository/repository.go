package repository

import (
	cDOM "github.com/Shakezidin/pkg/DOM/coordinator"
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
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

func (c *CoordinatorRepo) CreatePackage(pkg *cDOM.Package) error {
	if err := c.db.Create(&pkg).Error; err != nil {
		return err
	}
	return nil
}

func NewCoordinatorRepo(db *gorm.DB) inter.CoordinatorRepoInter {
	return &CoordinatorRepo{
		db: db,
	}
}

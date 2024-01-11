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

func NewCoordinatorRepo(db *gorm.DB) inter.CoordinatorRepoInter {
	return &CoordinatorRepo{
		db: db,
	}
}

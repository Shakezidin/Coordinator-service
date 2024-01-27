package repository

import (
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	"gorm.io/gorm"
)

type CoordinatorRepo struct {
	DB *gorm.DB
}

func NewCoordinatorRepo(db *gorm.DB) inter.CoordinatorRepoInter {
	return &CoordinatorRepo{
		DB: db,
	}
}

func (c *CoordinatorRepo) GetDB() *gorm.DB {
    return c.DB
}
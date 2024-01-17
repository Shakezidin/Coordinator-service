package repository

import (
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	"gorm.io/gorm"
)

type CoordinatorRepo struct {
	db *gorm.DB
}

func NewCoordinatorRepo(db *gorm.DB) inter.CoordinatorRepoInter {
	return &CoordinatorRepo{
		db: db,
	}
}

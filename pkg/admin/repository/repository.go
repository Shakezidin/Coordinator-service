package repository

import (
	DOM "github.com/Shakezidin/pkg/DOM/admin"
	inter "github.com/Shakezidin/pkg/admin/repository/interface"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) FetchAdmin(email string) (*DOM.Admin, error) {
	var admin DOM.Admin
	result := u.db.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func NewAdminRepository(db *gorm.DB) inter.RepoInterface {
	return &UserRepository{
		db: db,
	}
}

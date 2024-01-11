package admin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Password string	`gorm:"not null"`
	Role     string `gorm:"not null"`
}

package packages

import "gorm.io/gorm"

// User represents a user in the system.
type User struct {
	gorm.Model

	Name     string  `gorm:"not null"`
	Email    string  `gorm:"not null;unique"`
	Phone    int     `gorm:"not null;unique"`
	Password string  `gorm:"not null"`
	Role     string  `gorm:"not null"`
	Wallet   float64
	IsBlocked bool    `gorm:"default:false"` 

}

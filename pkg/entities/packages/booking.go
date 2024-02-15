package packages

import (
	"time"

	"gorm.io/gorm"
)

// ActivityBooking represents the booking of an activity by a traveller.
type ActivityBooking struct {
	gorm.Model

	TravellerID uint
	Traveller   Traveller `gorm:"foreignKey:TravellerID"`

	ActivityID uint
	Activity   Activity `gorm:"foreignKey:ActivityID"`
}

// Traveller represents a traveller who books activities and packages.
type Traveller struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       string
	Gender    string
	UserID    uint
	PackageID uint
	Package   Package `gorm:"foreignKey:PackageID"`

	Activities []Activity `gorm:"many2many:activity_booking;"`
}

// Booking represents a booking made by a user for a package.
type Booking struct {
	gorm.Model

	PaymentMode     string `gorm:"not null"`
	BookingStatus   string `gorm:"default:PENDING"`
	CancelledStatus string `gorm:"default:false"`
	PackagePrice    int
	PaidPrice       int
	UserID          uint
	UserEmail       string
	BookingID       string
	Bookings        []Traveller `gorm:"many2many:traveller_booking;"`
	PackageID       uint
	Package         Package `gorm:"foreignKey:PackageID"`
	BookDate        time.Time
	StartDate       time.Time
	CoordinatorID   uint
	CategoryID      uint
}

// RazorPay represents RazorPay payment details.
type RazorPay struct {
	UserID          uint
	RazorPaymentID  string
	RazorPayOrderID string
	Signature       string
	AmountPaid      float64
}

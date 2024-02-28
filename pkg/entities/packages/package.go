package packages

import (
	"time"

	"gorm.io/gorm"
)

// Category represents a trip category.
type Category struct {
	gorm.Model
	Category string `gorm:"not null;unique"`
}

// Package represents a tour package.
type Package struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	Destination      string    `gorm:"not null"`
	StartDate        time.Time `gorm:"not null"`
	StartTime        string    `gorm:"not null"`
	StartLocation    string    `gorm:"not null"`
	EndDate          time.Time `gorm:"not null"`
	MinPrice         int       `gorm:"not null"`
	MaxCapacity      int       `gorm:"not null"`
	NumOfDestination int       `gorm:"not null"`
	TripStatus       bool      `gorm:"default:false"`
	TripCategoryId   uint      `gorm:"not null"`
	Category         Category  `gorm:"foreignKey:TripCategoryId"`
	Images           string    `gorm:"not null"`
	CoordinatorID    uint      `gorm:"not null"`
	AvailableSpace   int       `gorm:"not null"`
	Description      string
}

// Destination represents a destination within a package.
type Destination struct {
	gorm.Model
	DestinationName    string `gorm:"not null"`
	Description        string `gorm:"not null"`
	PackageID          uint   `gorm:"not null"`
	Package            Package
	Image              string `gorm:"not null"`
	TransportationMode string `gorm:"not null"`
	ArrivalLocation    string `gorm:"not null"`
}

// Activity represents an activity within a destination.
type Activity struct {
	gorm.Model
	DestinationID uint `gorm:"not null"`
	Destination   Destination
	ActivityName  string `gorm:"not null"`
	Description   string `gorm:"not null"`
	Location      string `gorm:"not null"`
	ActivityType  string `gorm:"not null"`
	Amount        int    `gorm:"not null"`
	Date          time.Time
	Time          time.Time
}

// FoodMenu represents the food options for a package.
type FoodMenu struct {
	gorm.Model
	PackageID uint    `gorm:"not null"`
	Package   Package `gorm:"foreignKey:PackageID"`
	Breakfast string
	Lunch     string
	Dinner    string
	Date      string
}

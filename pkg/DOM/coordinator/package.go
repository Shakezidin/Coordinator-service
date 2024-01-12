package coordinator

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type JSONB []interface{}

// Value used to retrive value
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan helps to scan values
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type Category struct {
	gorm.Model
	Category string `gorm:"not null,unique"`
}

type Package struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	Destination      string    `gorm:"not null"`
	StartDate        time.Time `gorm:"not null"`
	StartLocation    string    `gorm:"not null"`
	EndDate          time.Time `gorm:"not null"`
	EndLoaction      string    `gorm:"not null"`
	Price            int       `gorm:"not null"`
	MaxCapacity      int       `gorm:"not null"`
	NumOfDestination int       `gorm:"not null"`
	TripStatus       bool      `gorm:"default:false"`
	TripCategoryId   uint      `gorm:"not null"`
	TripCategory     Category  `gorm:"ForeignKey:TripCategoryId"`
	Images           string    `gorm:"not null"`
	CoordinatorId    uint      `gorm:"not null"`
	Description      string
}

type Destination struct {
	gorm.Model
	DestinationName string  `gorm:"not null"`
	Description     string  `gorm:"not null"`
	PackageID       uint    `gorm:"not null"`
	Package         Package `gorm:"ForeignKey:PackageID"`
	MinPrice        int     `gorm:"not null"`
	MaxCapacity     int     `gorm:"not null"`
	Image           string
}

type Activity struct {
	gorm.Model
	DestinationId uint        `gorm:"not null"`
	Destination   Destination `gorm:"ForeignKey:DestinationId"`
	ActivityName  string      `gorm:"not null"`
	Description   string      `gorm:"not null"`
	Location      string      `gorm:"not null"`
	ActivityType  string      `gorm:"not null"`
	Amount        int         `gorm:"not null"`
	Date          time.Time
	time          time.Time
}

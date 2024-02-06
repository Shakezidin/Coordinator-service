package packages

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
	StartTime        string    `gorm:"not null"`
	StartLocation    string    `gorm:"not null"`
	EndDate          time.Time `gorm:"not null"`
	MinPrice         int       `gorm:"not null"`
	MaxCapacity      int       `gorm:"not null"`
	NumOfDestination int       `gorm:"not null"`
	TripStatus       bool      `gorm:"default:false"`
	TripCategoryId   uint      `gorm:"not null"`
	Category         Category  `gorm:"ForeignKey:TripCategoryId"`
	Images           string    `gorm:"not null"`
	CoordinatorId    uint      `gorm:"not null"`
	Availablespace   int       `grom:"not null"`
	Description      string
}

type Destination struct {
	gorm.Model
	DestinationName    string  `gorm:"not null"`
	Description        string  `gorm:"not null"`
	PackageID          uint    `gorm:"not null"`
	Package            Package `gorm:"ForeignKey:PackageID"`
	Image              string
	TransportationMode string `gorm:"not null"`
	ArrivalLocation    string `gorm:"not null"`
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
	Time          time.Time
}

type ActivityBooking struct {
	gorm.Model
	TravellerId uint
	Traveller   Traveller `gorm:"foreignKey:TravellerId"`
	ActivityId  uint
	Activity    Activity `gorm:"foreignKey:ActivityId"`
}

type Traveller struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       string
	Gender    string
	UserId    uint
	PackageId uint
	Package   Package    `gorm:"foreignKey:PackageId"`
	Activity  []Activity `gorm:"many2many:activity_booking;"`
}

type Booking struct {
	gorm.Model
	PaymentMode     string `gorm:"not null"`
	BookingStatus   string `gorm:"default:PENDING"`
	CancelledStatus string `gorm:"default:false"`
	TotalPrice      int
	UserId          uint
	BookingId       string
	Bookings        []Traveller `gorm:"many2many:traveller_booking;"`
	PackageId       uint
	Package         Package `gorm:"foreignKey:PackageId"`
	BookDate        time.Time
	StartDate       time.Time
	CoordinatorID   uint
}

type RazorPay struct {
	UserID          uint
	RazorPaymentID  string
	RazorPayOrderID string
	Signature       string
	AmountPaid      float64
}

type FoodMenu struct {
	gorm.Model
	PackageId uint
	Package   Package `gorm:"foreignKey:PackageId"`
	Breakfast string
	Lunch     string
	Dinner    string
	Date      string
}

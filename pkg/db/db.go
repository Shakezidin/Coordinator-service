package db

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	cDOM "github.com/Shakezidin/pkg/entities/packages"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database(config *config.Config) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to the database failed:", err)
	}

	// AutoMigrate all models
	err = DB.AutoMigrate(
		cDOM.Package{},
		cDOM.Category{},
		cDOM.Activity{},
		cDOM.Destination{},
		cDOM.User{},
		cDOM.ActivityBooking{},
		cDOM.Traveller{},
		cDOM.Booking{},
	)
	if err != nil {
		fmt.Println("error while migrating")
		return nil
	}

	return DB
}

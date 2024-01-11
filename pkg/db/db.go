package db

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	DOM "github.com/Shakezidin/pkg/DOM/admin"
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
	DB.AutoMigrate(DOM.Admin{})

	return DB
}
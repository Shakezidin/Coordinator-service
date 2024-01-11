package di

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/admin/handler"
	"github.com/Shakezidin/pkg/admin/service"
	"github.com/Shakezidin/pkg/db"
	"github.com/Shakezidin/pkg/admin/server"
	"github.com/Shakezidin/pkg/admin/repository"
)

func Init() {
	config := config.LoadConfig()
	db := db.Database(config)
	adminrepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminrepo)
	adminHandler := handler.NewAdminHandler(adminService)
	err := server.NewAdminGrpcServer(config, adminHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}

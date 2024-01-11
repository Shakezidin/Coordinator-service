package di

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/admin/handler"
	"github.com/Shakezidin/pkg/admin/server"
	"github.com/Shakezidin/pkg/admin/service"
	"github.com/Shakezidin/pkg/db"
	"github.com/Shakezidin/pkg/server"
)

func Init() {
	config := config.LoadConfig()
	db := db.Database(config)
	client, err :=server.NewAdminGrpcServer()
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
	adminrepo := repository.AdminRepository(db)
	adminService := service.AdminService(adminrepo, client)
	adminHandler := handler.AdminHandler(adminService, config)
	err = server.NewGrpcServer(config, adminHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}

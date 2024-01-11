package di

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/coordinator/handler"
	"github.com/Shakezidin/pkg/coordinator/repository"
	"github.com/Shakezidin/pkg/coordinator/server"
	"github.com/Shakezidin/pkg/coordinator/service"
	"github.com/Shakezidin/pkg/db"
)

func Init() {
	config := config.LoadConfig()
	db := db.Database(config)
	coordinatorepo := repository.NewCoordinatorRepo(db)
	coordinatorService := service.NewCoordinatorSVC(coordinatorepo)
	coordinatorHandler := handler.NewCoordinatorHandler(coordinatorService)
	err := server.NewCoordinatorGrpcServer(config, coordinatorHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}

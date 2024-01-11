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
	cnfg := config.LoadConfig()
	redis:=config.ConnectToRedis(cnfg)
	twilio:=config.SetupTwilio(cnfg)
	db := db.Database(cnfg)
	coordinatorepo := repository.NewCoordinatorRepo(db)
	coordinatorService := service.NewCoordinatorSVC(coordinatorepo,twilio,redis,cnfg)
	coordinatorHandler := handler.NewCoordinatorHandler(coordinatorService)
	err := server.NewCoordinatorGrpcServer(cnfg, coordinatorHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}

package di

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	client "github.com/Shakezidin/pkg/coordinator/client"
	"github.com/Shakezidin/pkg/coordinator/handler"
	"github.com/Shakezidin/pkg/coordinator/repository"
	"github.com/Shakezidin/pkg/coordinator/server"
	"github.com/Shakezidin/pkg/coordinator/service"
	"github.com/Shakezidin/pkg/db"
)

func Init() {
	cnfg := config.LoadConfig()
	redis := config.ConnectToRedis(cnfg)
	twilio := config.SetupTwilio(cnfg)
	db := db.Database(cnfg)
	client, err := client.ClientDial(*cnfg)
	if err != nil {
		fmt.Println("client dial error")
		return
	}
	coordinatorepo := repository.NewCoordinatorRepo(db)
	coordinatorService := service.NewCoordinatorSVC(coordinatorepo, twilio, redis, cnfg, client)
	coordinatorHandler := handler.NewCoordinatorHandler(coordinatorService)
	config.InitCron(coordinatorService)
	err = server.NewCoordinatorGrpcServer(cnfg, coordinatorHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}

package di

import (
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

	// Connect to Redis
	redis, err := config.ConnectToRedis(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	// Setup Twilio client
	twilio := config.SetupTwilio(cnfg)

	// Connect to the database
	db := db.Database(cnfg)

	// Dial gRPC client
	client, err := client.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to dial gRPC client: %v", err)
	}

	// Initialize repository
	coordinatorepo := repository.NewCoordinatorRepo(db)

	// Initialize coordinator service
	coordinatorService := service.NewCoordinatorSVC(coordinatorepo, twilio, redis, cnfg, client)

	// Initialize coordinator handler
	coordinatorHandler := handler.NewCoordinatorHandler(coordinatorService)

	// Initialize cron jobs
	config.InitCron(coordinatorService)

	// Start gRPC server
	err = server.NewCoordinatorGrpcServer(cnfg, coordinatorHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}

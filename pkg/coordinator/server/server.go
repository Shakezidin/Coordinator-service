package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/coordinator/handler"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"google.golang.org/grpc"
)

func NewCoordinatorGrpcServer(cfg *config.Config, handlr *handler.CoordinatorHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.GRPCCOORDINATORPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	cpb.RegisterCoordinatorServer(grp, handlr)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", cfg.GRPCCOORDINATORPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}

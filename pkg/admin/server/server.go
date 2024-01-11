package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Shakezidin/config"
	"google.golang.org/grpc"
	"github.com/Shakezidin/pkg/admin/handler"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

func NewAdminGrpcServer(cfg *config.Config, handlr *handler.AdminHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.GRPCADMINPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	adminpb.RegisterAdminServer(grp,handlr)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", cfg.GRPCADMINPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}

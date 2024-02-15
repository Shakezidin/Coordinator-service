package client

import (
	"log"

	"github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/coordinator/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.AdminClient, error) {
	grpc, err := grpc.Dial(cfg.GRPCADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCADMINPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to coordinator Client at port: %v", cfg.GRPCADMINPORT)
	return pb.NewAdminClient(grpc), nil
}

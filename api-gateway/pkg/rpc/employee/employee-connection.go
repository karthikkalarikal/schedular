package employee

import (
	"log"

	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConnection struct {
	Connection *grpc.ClientConn
}

func InitRPC(cfg config.Config) *GRPCConnection {
	conn, err := grpc.Dial(cfg.EmployeeGRPC, grpc.WithTransportCredentials(insecure.NewCredentials())) //not secure

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)

	}
	defer conn.Close()

	return &GRPCConnection{
		Connection: conn,
	}

}

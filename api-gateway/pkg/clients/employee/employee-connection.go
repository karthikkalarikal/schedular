package employee

import (
	"log"

	"github.com/karthikkalarikal/api-gateway/pkg/clients/employee/proto"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EmployeeClient struct {
	Client proto.EmployeeServiceClient
}

func NewEmployeeClient(cfg config.Config) *EmployeeClient {
	conn, err := grpc.Dial(cfg.EmployeeGRPC, grpc.WithTransportCredentials(insecure.NewCredentials())) //not secure

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)

	}
	defer conn.Close()

	return &EmployeeClient{
		Client: proto.NewEmployeeServiceClient(conn),
	}

}

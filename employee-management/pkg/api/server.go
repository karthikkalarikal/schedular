package HTTP

import (
	"fmt"
	"log"
	"net"

	"github.com/karthikkalarikal/employee-management/pkg/api/service"
	"github.com/karthikkalarikal/employee-management/pkg/config"
	"github.com/karthikkalarikal/employee-management/pkg/pb"

	"google.golang.org/grpc"
)

type ServiceHTTP struct {
	grpc *grpc.Server
	port string
}

func NewServerHTTP(config config.Config, employee *service.EmployeeService) *ServiceHTTP {
	s := grpc.NewServer()

	pb.RegisterEmployeeServiceServer(s, employee)

	return &ServiceHTTP{
		grpc: s,
		port: config.Port,
	}
}

func (s *ServiceHTTP) Start() {
	fmt.Println("tcp", s.port)
	lis, err := net.Listen("tcp", s.port)

	if err != nil {
		log.Fatalf("cannot listen to port: %v", err)
	}
	if err = s.grpc.Serve(lis); err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}

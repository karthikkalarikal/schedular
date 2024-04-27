package HTTP

import (
	"log"
	"net"

	"github.com/karthikkalarikal/employee-management/pkg/config"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpc *grpc.Server
	port string
}

func InitGRPC(config config.Config) *GrpcServer {
	s := grpc.NewServer()

	return &GrpcServer{
		grpc: s,
		port: config.Port,
	}
}

func (s *GrpcServer) Start() {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalf("cannot listen to port: %v", err)
	}
	if err = s.grpc.Serve(lis); err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}

package service

import (
	"context"
	"fmt"

	"github.com/karthikkalarikal/employee-management/pkg/pb"
	"github.com/karthikkalarikal/employee-management/pkg/repository/interfaces"
)

type EmployeeService struct {
	repo interfaces.AdminRepo
	pb.UnimplementedEmployeeServiceServer
}

func NewEmployeeService(repo interfaces.AdminRepo) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

func (s *EmployeeService) AddEmployee(ctx context.Context, admin *pb.Employee) (*pb.EmployeeReply, error) {
	fmt.Println("here in service")
	response, err := s.repo.AddEmployee(ctx, admin)
	if err != nil {
		return nil, err
	}

	return response, nil

}

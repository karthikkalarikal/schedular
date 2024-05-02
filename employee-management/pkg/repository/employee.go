package repository

import (
	"context"

	"github.com/karthikkalarikal/employee-management/pkg/pb"
	"github.com/karthikkalarikal/employee-management/pkg/repository/interfaces"
)

type adminRepoImpl struct {
}

func NewAdminRepo() interfaces.AdminRepo {
	return &adminRepoImpl{}
}

func (a *adminRepoImpl) AddEmployee(ctx context.Context, res *pb.Employee) (*pb.EmployeeReply, error) {

	return &pb.EmployeeReply{
		Id:         1,
		Name:       "name",
		Department: "department",
		Role:       "role",
		CreatedAt:  "now()",
	}, nil
}

package interfaces

import (
	"context"

	"github.com/karthikkalarikal/employee-management/pkg/pb"
)

type AdminRepo interface {
	AddEmployee(context.Context, *pb.Employee) (*pb.EmployeeReply, error)
}

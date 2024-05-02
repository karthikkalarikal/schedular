package usecase

import (
	"fmt"

	"github.com/karthikkalarikal/api-gateway/pkg/clients/employee"
	"github.com/karthikkalarikal/api-gateway/pkg/clients/employee/proto"
	"github.com/karthikkalarikal/api-gateway/pkg/domain"
	"github.com/karthikkalarikal/api-gateway/pkg/usecase/interfaces"
	"github.com/labstack/echo/v4"
)

type administratorUsecaseImpl struct {
	client *employee.EmployeeClient
}

func NewAdministratorUsecase(client *employee.EmployeeClient) interfaces.AdministratorUsecaseInterfaces {
	return &administratorUsecaseImpl{
		client: client,
	}
}

func (u *administratorUsecaseImpl) AddEmployee(e echo.Context, employee domain.Employee) (res domain.EmployeeResponse, err error) {
	// res,er
	fmt.Println("here in usecase")
	body, err := u.client.Client.AddEmployee(e.Request().Context(), &proto.Employee{
		Name:       employee.Name,
		Department: employee.Department,
		Role:       employee.Role,
	})
	fmt.Println("err", err)
	if err != nil {
		return domain.EmployeeResponse{}, err
	}
	fmt.Println("err", err)

	return domain.EmployeeResponse{
		Id:         int(body.Id),
		Name:       body.Name,
		Department: body.Department,
		Role:       body.Role,
		CreatedAt:  body.CreatedAt,
	}, nil
}

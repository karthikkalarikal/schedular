package usecase

import (
	"time"

	"github.com/karthikkalarikal/api-gateway/pkg/domain"
	"github.com/karthikkalarikal/api-gateway/pkg/usecase/interfaces"
)

type administratorUsecaseImpl struct {
}

func NewAdministratorUsecase() interfaces.AdministratorUsecaseInterfaces {
	return &administratorUsecaseImpl{}
}

func (u *administratorUsecaseImpl) AddEmployee(employee domain.Employee) (res domain.EmployeeResponse, err error) {
	// res,er
	return domain.EmployeeResponse{
		Id:         1,
		Name:       "name-1",
		Department: "dep-1",
		CreatedAt:  time.Now(),
	}, nil
}

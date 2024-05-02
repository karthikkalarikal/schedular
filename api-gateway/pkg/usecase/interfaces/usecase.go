package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/domain"
	"github.com/labstack/echo/v4"
)

type AdministratorUsecaseInterfaces interface {
	AddEmployee(echo.Context, domain.Employee) (domain.EmployeeResponse, error)
}

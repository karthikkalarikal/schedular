package interfaces

import "github.com/karthikkalarikal/api-gateway/pkg/domain"

type AdministratorUsecaseInterfaces interface {
	AddEmployee(domain.Employee) (domain.EmployeeResponse, error)
}

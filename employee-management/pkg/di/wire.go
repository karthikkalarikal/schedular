package di

import (
	"log"

	HTTP "github.com/karthikkalarikal/employee-management/pkg/api"
	"github.com/karthikkalarikal/employee-management/pkg/api/service"
	"github.com/karthikkalarikal/employee-management/pkg/config"
	"github.com/karthikkalarikal/employee-management/pkg/repository"
)

func InitializeAPI() *HTTP.ServiceHTTP {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error in wire %v", err)
	}

	repo := repository.NewAdminRepo()
	empSer := service.NewEmployeeService(repo)
	serv := HTTP.NewServerHTTP(c, empSer)
	return serv
}

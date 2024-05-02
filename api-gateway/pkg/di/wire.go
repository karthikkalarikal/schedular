package di

import (
	"github.com/karthikkalarikal/api-gateway/pkg/api"
	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers"
	"github.com/karthikkalarikal/api-gateway/pkg/clients/employee"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/usecase"
)

func InitializeAPI() (*api.Server, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	empClient := employee.NewEmployeeClient(cfg)
	adUse := usecase.NewAdministratorUsecase(empClient)
	adHand := handlers.NewHandler(adUse)
	server := api.NewServerHTTP(cfg, *adHand)

	return server, nil
}

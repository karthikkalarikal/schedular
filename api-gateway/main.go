package main

import (
	"log"

	_ "github.com/karthikkalarikal/api-gateway/docs"
	"github.com/karthikkalarikal/api-gateway/pkg/di"
)

func main() {

	service, err := di.InitializeAPI()
	if err != nil {
		log.Fatalf("failed to initialize api, error: %s", err.Error())
	}

	service.Start()
}

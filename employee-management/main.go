package main

import (
	"github.com/karthikkalarikal/employee-management/pkg/di"
)

func main() {
	service := di.InitializeAPI()

	service.Start()
}

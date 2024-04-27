package routes

import (
	"fmt"

	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo, administrator handlers.AdministratorHandler) {
	fmt.Println("here")
	employee := e.Group("/employee")

	{
		employee.POST("/", administrator.AddEmployee)
		// employee.PUT("/:id", administrator.Modify)
	}
}

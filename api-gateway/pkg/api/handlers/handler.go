package handlers

import (
	"fmt"
	"net/http"

	"github.com/karthikkalarikal/api-gateway/pkg/domain"
	"github.com/karthikkalarikal/api-gateway/pkg/usecase/interfaces"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AdministratorHandler struct {
	usecase interfaces.AdministratorUsecaseInterfaces
}

func NewHandler(usecase interfaces.AdministratorUsecaseInterfaces) *AdministratorHandler {
	return &AdministratorHandler{
		usecase: usecase,
	}
}

// Administrator godoc
//
//	@Summary		Add an employee
//	@Description	This function is to add an employee to the data base
//	@Tags			Administration
//	@Accept			json
//	@Produce		json
//	@Param			employee	body		domain.Employee	true	"Employee details"
//	@Success		200			{object}	domain.Response
//	@Failure		400			{object}	domain.Response
//	@Failure		401			{object}	domain.Response
//	@Failure		404			{object}	domain.Response
//	@Failure		500			{object}	domain.Response
//	@Router			/employee/ [post]
func (h *AdministratorHandler) AddEmployee(e echo.Context) error {
	fmt.Println("insert employee administrator")

	var employee domain.Employee

	if err := e.Bind(&employee); err != nil {
		return e.JSONPretty(http.StatusBadRequest, echo.Map{"error": err.Error()}, " ")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(employee); err != nil {
		return e.JSONPretty(http.StatusBadRequest, echo.Map{"error": err.Error()}, " ")
	}

	response, err := h.usecase.AddEmployee(e, employee)
	fmt.Println("response", response)
	if err != nil {
		return e.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err.Error()}, " ")
	}
	return e.JSONPretty(http.StatusOK, response, " ")
}

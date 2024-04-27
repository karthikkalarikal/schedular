package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	engine *echo.Echo
	port   string
}

func NewServerHTTP(cfg config.Config, administrator handlers.AdministratorHandler) *Server {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodConnect},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	routes.SetUpRoutes(e, administrator)

	return &Server{
		engine: e,
		port:   cfg.Port,
	}
}

func (s *Server) Start() {
	fmt.Println("port: ", s.port)

	if err := s.engine.Start(fmt.Sprintf(":%s", s.port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

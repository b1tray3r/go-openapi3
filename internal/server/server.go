package server

import (
	"context"
	"net/http"

	"github.com/b1tray3r/go-openapi3/pkg/api"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func (s Server) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	message := "Service is healthy"
	response := api.GetHealth200JSONResponse{
		HealthCheckResponseJSONResponse: api.HealthCheckResponseJSONResponse{
			Message: &message,
		},
	}

	return response, nil
}

func NewEchoServer() *echo.Echo {
	e := echo.New()

	api.RegisterHandlers(e, api.NewStrictHandler(
		Server{},
		// add middlewares here if needed
		[]api.StrictMiddlewareFunc{},
	))

	return e
}

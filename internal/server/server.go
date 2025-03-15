package server

import (
	"context"

	"github.com/b1tray3r/go-openapi3/pkg/api"
	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

type Server struct{}

func (s Server) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	message := "Service is healthy"
	response := api.GetHealth200JSONResponse{
		SuccessResponseJSONResponse: api.SuccessResponseJSONResponse{
			Message: &message,
		},
	}

	return response, nil
}

func (s Server) GetTest(ctx context.Context, request api.GetTestRequestObject) (api.GetTestResponseObject, error) {
	message := "Test endpoint"
	response := api.GetTest200JSONResponse{
		SuccessResponseJSONResponse: api.SuccessResponseJSONResponse{
			Message: &message,
		},
	}

	return response, nil
}

func StrictTokenAuthMiddleware(f strictecho.StrictEchoHandlerFunc, operationID string) strictecho.StrictEchoHandlerFunc {
	return func(ctx echo.Context, request interface{}) (response interface{}, err error) {
		auth := &Auth{
			Secret: "123456",
			PublicEndpoints: []string{
				"/health",
			},
		}
		httpError := auth.Authenticate(ctx)
		if httpError != nil {
			return nil, httpError
		}

		return f(ctx, request)
	}
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	api.RegisterHandlers(e, api.NewStrictHandler(
		Server{},
		[]api.StrictMiddlewareFunc{
			StrictTokenAuthMiddleware,
		},
	))

	return e
}

package openapi

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

func NewDefinition() *openapi3.T {
	return &openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       "OpenAPI3 Example API",
			Description: "API to demo a REST API definition with OpenAPI in swagger UI.",
			Version:     "1.0.0",
			Contact: &openapi3.Contact{
				Name:  "aborgardt",
				URL:   "https://github.com/b1tray3r/",
				Email: "5030347+b1tray3r@users.noreply.github.com",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "localhost",
				URL:         "http://localhost:8085",
			},
		},
		Paths: openapi3.NewPaths(
			openapi3.WithPath(
				"/health",
				&openapi3.PathItem{
					Summary:     "HealthCheck without authentication",
					Description: "Checks the availablility of the API server.",
					Get: &openapi3.Operation{
						Tags:    []string{"Public"},
						Summary: "Get health status",
						Responses: openapi3.NewResponses(
							openapi3.WithStatus(
								http.StatusOK,
								&openapi3.ResponseRef{
									Ref: "#/components/responses/SuccessResponse",
								},
							),
							openapi3.WithStatus(
								http.StatusInternalServerError,
								&openapi3.ResponseRef{
									Ref: "#/components/responses/ErrorResponse",
								},
							),
						),
					},
				},
			),
			openapi3.WithPath(
				"/test",
				&openapi3.PathItem{
					Summary:     "Test endpoint with authentication",
					Description: "Checks the availablility of the API server.",
					Get: &openapi3.Operation{
						Tags:    []string{"Private"},
						Summary: "Test endpoint with authentication",
						Security: openapi3.NewSecurityRequirements().
							With(openapi3.NewSecurityRequirement().Authenticate("apiKeyAuth")),
						Responses: openapi3.NewResponses(
							openapi3.WithStatus(
								http.StatusOK,
								&openapi3.ResponseRef{
									Ref: "#/components/responses/SuccessResponse",
								},
							),
							openapi3.WithStatus(
								http.StatusUnauthorized,
								&openapi3.ResponseRef{
									Ref: "#/components/responses/ErrorResponse",
								},
							),
							openapi3.WithStatus(
								http.StatusInternalServerError,
								&openapi3.ResponseRef{
									Ref: "#/components/responses/ErrorResponse",
								},
							),
						),
					},
				},
			),
		),
		Components: &openapi3.Components{
			SecuritySchemes: openapi3.SecuritySchemes{
				"apiKeyAuth": &openapi3.SecuritySchemeRef{
					Value: openapi3.NewSecurityScheme().
						WithDescription("API Key authentication").
						WithType("apiKey").
						WithName("Authorization").
						WithIn("header").
						WithBearerFormat("bearer"),
				},
			},
			Schemas: openapi3.Schemas{
				"Status": openapi3.NewSchemaRef("",
					openapi3.NewObjectSchema().
						WithProperty(
							"message",
							openapi3.NewStringSchema().WithDefault("Some message from the backend!"),
						),
				),
			},
			Responses: map[string]*openapi3.ResponseRef{
				"ErrorResponse": {
					Value: openapi3.NewResponse().
						WithDescription("Something went wrong!").
						WithContent(
							openapi3.NewContentWithJSONSchemaRef(
								&openapi3.SchemaRef{
									Ref: "#/components/schemas/Status",
								},
							),
						),
				},
				"SuccessResponse": {
					Value: openapi3.NewResponse().
						WithDescription("Server is running!").
						WithContent(
							openapi3.NewContentWithJSONSchemaRef(
								&openapi3.SchemaRef{
									Ref: "#/components/schemas/Status",
								},
							),
						),
				},
			},
		},
	}
}

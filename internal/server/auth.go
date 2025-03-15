package server

import (
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	PublicEndpoints []string
	Secret          string
}

func (a *Auth) Authenticate(ctx echo.Context) *echo.HTTPError {
	if slices.Contains(a.PublicEndpoints, ctx.Request().RequestURI) {
		return nil
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
	}

	token := strings.TrimPrefix(authHeader, "bearer ")
	if token != a.Secret {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
	return nil
}

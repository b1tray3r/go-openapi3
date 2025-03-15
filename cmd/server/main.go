package main

import (
	"embed"
	"net/http"

	"github.com/b1tray3r/go-openapi3/internal/server"
	"github.com/labstack/echo/v4"
)

//go:embed swagger/*
var StaticFiles embed.FS

func main() {
	srv := server.NewEchoServer()

	srv.GET(
		"/swagger/*",
		echo.WrapHandler(
			http.StripPrefix(
				"/",
				http.FileServer(
					http.FS(
						StaticFiles,
					),
				),
			),
		),
	)

	if err := srv.Start("localhost:8085"); err != nil {
		panic(err.Error())
	}
}

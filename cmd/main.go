package main

import (
	"github.com/b1tray3r/go-openapi3/internal/server"
)

func main() {
	srv := server.NewEchoServer()

	if err := srv.Start("localhost:8085"); err != nil {
		panic(err.Error())
	}
}

package main

import (
	"github.com/lus/ws/server"
	"net/http"
)

func main() {
	settings := &server.Settings{
		Address:    "0.0.0.0",
		Port:       1337,
		FileSystem: http.Dir("."),
	}
	panic(server.Start(settings))
}

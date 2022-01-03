package server

import (
	"fmt"
	"net/http"
)

// Settings represents the settings the server starts with
type Settings struct {
	Address    string
	Port       int
	FileSystem http.FileSystem
}

// Start starts the server with the specified settings
func Start(settings *Settings) error {
	address := fmt.Sprintf("%s:%d", settings.Address, settings.Port)
	return http.ListenAndServe(address, http.FileServer(settings.FileSystem))
}

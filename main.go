package main

import (
	"flag"
	"fmt"
	"github.com/lus/ws/out"
	"github.com/lus/ws/server"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	// Parse command line flags
	address := flag.String("a", "0.0.0.0", "the address to listen to")
	port := flag.Int("p", 1337, "the port to listen to")
	noFormatting := flag.Bool("nf", false, "whether to disable output formatting")
	flag.Parse()

	// Potentially disable output formatting
	if *noFormatting {
		out.DisableFormatting()
	}

	// Validate the specified path and turn it into a clean absolute representation
	rawPath := "."
	if len(os.Args) > 1 {
		argPath := os.Args[len(os.Args)-1]
		if !strings.HasPrefix(argPath, "-") {
			rawPath = argPath
		}
	}
	absPath, err := filepath.Abs(rawPath)
	if err != nil {
		out.Error(err.Error())
	}
	stat, err := os.Stat(absPath)
	if err != nil {
		out.Error(err.Error())
	}
	if !stat.IsDir() {
		out.Error("the given filepath points to a file and not a directory")
	}
	out.Info("compiled final path: " + absPath)

	// Build the settings and start the server
	settings := &server.Settings{
		Address:    *address,
		Port:       *port,
		FileSystem: http.Dir(absPath),
	}
	out.Info(fmt.Sprintf("starting web server on %s:%d", settings.Address, settings.Port))
	go func() {
		if err := server.Start(settings); err != nil {
			out.Error(err.Error())
		}
	}()

	// Wait for the program to exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-sc
	out.Info("stopping the web server")
}

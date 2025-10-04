package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/kiriyms/server_devops_practice_go/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	mux.Handle("/", Logger(handlers.NewVisitorHandler()))

	addr := "0.0.0.0:8080"
	logger.Info("Server is running", slog.String("address", "http://"+addr))
	http.ListenAndServe(addr, mux)
}

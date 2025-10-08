package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/kiriyms/server_devops_practice_go/common"
	"github.com/kiriyms/server_devops_practice_go/handlers"
	"github.com/kiriyms/server_devops_practice_go/services"
)

func main() {
	common.MustLoadConfig()
	common.LoadLogger()
	cfg := common.GetConfig()

	slog.Info("Config loaded.")
	slog.Info("Logger loaded.", slog.String("env", cfg.Environment))
	slog.Debug("Debug logs enabled.")

	service := services.NewGreeter()
	loggedService := services.NewLoggingService(service)

	handler := handlers.NewHandler(loggedService)
	loggedHandler := handlers.NewLoggingHandler(handler)

	mux := http.NewServeMux()
	mux.Handle("/", loggedHandler)

	addr := fmt.Sprintf("%s:%s", cfg.Address, cfg.Port)
	slog.Info("Server is running.", slog.String("addr", "http://"+addr))
	http.ListenAndServe(addr, mux)
}

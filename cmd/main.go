package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/kiriyms/server_devops_practice_go/common"
	"github.com/kiriyms/server_devops_practice_go/handlers"
)

func main() {
	config := common.MustLoadConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	mux.Handle("/", Logger(handlers.NewVisitorHandler()))

	addr := fmt.Sprintf("%s:%s", config.Address, config.Port)
	logger.Info("Server is running", slog.String("address", "http://"+addr))
	http.ListenAndServe(addr, mux)
}

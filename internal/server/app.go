// Package server is used to run app.
package server

import (
	"fmt"
	"net/http"

	"github.com/kozyrev-m/effective-mobile-task/internal/logger"
	httpserver "github.com/kozyrev-m/effective-mobile-task/internal/server/http-server"
	"go.uber.org/zap"
)

var (
	addr  = ":8080"
	level = zap.DebugLevel.String()
)

// StartApp launches the application.
func StartApp() error {
	if err := logger.InitLogger(level); err != nil {
		return err
	}

	srv := &http.Server{
		Handler: httpserver.New(),
		Addr:    addr,
	}

	logger.Log.Info(fmt.Sprintf("Running server with log level '%s'", level), zap.String("address", addr))

	return srv.ListenAndServe()
}

// Package server is used to run app.
package server

import (
	"fmt"
	"net/http"

	"github.com/kozyrev-m/effective-mobile-task/internal/config"
	"github.com/kozyrev-m/effective-mobile-task/internal/logger"
	httpserver "github.com/kozyrev-m/effective-mobile-task/internal/server/http-server"
	"github.com/kozyrev-m/effective-mobile-task/internal/store/pg"
	"go.uber.org/zap"
)

// StartApp launches the application.
func StartApp(cfg *config.Config) error {
	if err := logger.InitLogger(cfg.LogLevel); err != nil {
		return err
	}

	store, err := pg.NewStore(cfg.DatabaseDSN)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Handler: httpserver.New(store),
		Addr:    cfg.Address,
	}

	logger.Log.Info(fmt.Sprintf("Running server with log level '%s'", cfg.LogLevel), zap.String("address", cfg.Address))

	return srv.ListenAndServe()
}

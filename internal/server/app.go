// Package server is used to run app.
package server

import (
	"net/http"

	httpserver "github.com/kozyrev-m/effective-mobile-task/internal/server/http-server"
)

const addr = ":8080"

// StartApp launches the application.
func StartApp() error {
	srv := &http.Server{
		Handler: httpserver.New(),
		Addr:    addr,
	}

	return srv.ListenAndServe()
}

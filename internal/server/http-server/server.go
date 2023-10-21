// Package httpserver implements http-server, configures handlers, middleware and more.
package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPServer represents http-server.
type HTTPServer struct {
	router *gin.Engine
}

// New is a constructor to create HTTPServer.
func New() *HTTPServer {
	return &HTTPServer{
		router: gin.New(),
	}
}

// ServeHTTP implements http.Handler interface.
// Note: for compatibility with the net/http package.
func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

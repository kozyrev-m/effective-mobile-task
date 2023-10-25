// Package httpserver implements http-server, configures handlers, middleware and more.
package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kozyrev-m/effective-mobile-task/internal/integration"
	"github.com/kozyrev-m/effective-mobile-task/internal/service"
	"github.com/kozyrev-m/effective-mobile-task/internal/store"
)

// HTTPServer represents http-server.
type HTTPServer struct {
	router  *gin.Engine
	service *service.Service
	agent   integration.Integration
}

// New is a constructor to create HTTPServer.
func New(store store.Store, agent integration.Integration) *HTTPServer {
	gin.SetMode(gin.ReleaseMode)

	svc := service.NewService(store)

	s := &HTTPServer{
		router:  gin.New(),
		service: svc,
		agent:   agent,
	}

	s.initRouter()

	return s
}

// initRouter sets endpoints.
func (s *HTTPServer) initRouter() {
	s.router.Use(s.dummyMiddleware())
	s.router.Use(s.setRequestID())
	s.router.Use(s.withLogging())

	s.router.GET("/find/:id", s.handlerFindPerson)
	s.router.DELETE("/delete/:id", s.handlerDeletePerson)
	s.router.PATCH("/update/:id", s.handlerUpdatePerson)
	s.router.POST("/add", s.handlerAddPerson)
	s.router.POST("/persons", s.handlerPersons)

}

// ServeHTTP implements http.Handler interface.
// Note: for compatibility with the net/http package.
func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

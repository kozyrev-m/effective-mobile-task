package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kozyrev-m/effective-mobile-task/internal/logger"
	"go.uber.org/zap"
)

// dummyMiddleware is used as a test middleware.
func (s *HTTPServer) dummyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// logger.Log.Info("into dummyMiddleware")

		ctx.Next()
	}
}

// withLogging is used to log request and response information for all endpoints.
func (s *HTTPServer) withLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		clientIP := ctx.ClientIP()

		logger.Log.Info(
			"incoming request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.RequestURI),
			zap.String("user_ip", clientIP),
		)

		ctx.Next()

		code := ctx.Writer.Status()

		logger.Log.Info(
			"completed request",
			zap.Int("status_code", code),
			zap.String("status_text", http.StatusText(code)),
			zap.String("user_ip", clientIP),
			zap.Reflect("duration", time.Since(start)),
		)

		ctx.Next()
	}
}

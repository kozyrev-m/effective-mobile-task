package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kozyrev-m/effective-mobile-task/internal/logger"
	"go.uber.org/zap"
)

const (
	keyRequestID string = "request-id"
)

// dummyMiddleware is used as a test middleware.
func (s *HTTPServer) dummyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// logger.Log.Info("into dummyMiddleware")

		ctx.Next()
	}
}

// setRequestID creates and sets request id.
func (s *HTTPServer) setRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.New().String()
		ctx.Writer.Header().Set("X-Request-Id", id)
		ctx.Set(keyRequestID, id)

		ctx.Next()
	}
}

// withLogging is used to log request and response information for all endpoints.
func (s *HTTPServer) withLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// determine request id
		var requestID string = ""
		ctxRequestID, exists := ctx.Get(keyRequestID)
		if exists {
			requestID = ctxRequestID.(string)
		}

		// customize logger
		customLogger := logger.Log.WithOptions(
			zap.Fields(
				zap.String("request_id", requestID),
				zap.String("user_ip", ctx.ClientIP()),
			),
		)

		// after request
		customLogger.Info(
			"incoming request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.RequestURI),
		)

		ctx.Next()

		// after response
		code := ctx.Writer.Status()
		customLogger.Info(
			"completed request",
			zap.Int("status_code", code),
			zap.String("status_text", http.StatusText(code)),
			zap.Reflect("duration", time.Since(start)),
		)

		ctx.Next()
	}
}

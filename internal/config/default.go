package config

import "go.uber.org/zap"

// default configs.
var (
	// address and port to run server.
	defaultRunAddr = "127.0.0.1:8080"
	// log level.
	defaultLogLevel = zap.DebugLevel.CapitalString()
)

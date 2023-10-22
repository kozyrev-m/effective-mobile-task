// Package config works with app config.
package config

import (
	"flag"
	"os"
)

var (
	flagRunAddr  string
	flagLogLevel string
)

// Config contains app configuration.
type Config struct {
	Address  string
	LogLevel string
}

// InitConfig sets app config.
func InitConfig() *Config {
	flag.StringVar(&flagRunAddr, "a", defaultRunAddr, "address and port to run server")
	flag.StringVar(&flagLogLevel, "l", defaultLogLevel, "log level")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagLogLevel = envLogLevel
	}

	config := &Config{
		Address:  flagRunAddr,
		LogLevel: flagLogLevel,
	}

	return config
}

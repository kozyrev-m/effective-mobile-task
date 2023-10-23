// Package config works with app config.
package config

import (
	"flag"
	"os"
)

var (
	flagRunAddr     string
	flagLogLevel    string
	flagDatabaseDSN string
)

// Config contains app configuration.
type Config struct {
	Address     string
	LogLevel    string
	DatabaseDSN string
}

// InitConfig sets app config.
func InitConfig() *Config {
	flag.StringVar(&flagRunAddr, "a", defaultRunAddr, "address and port to run server")
	flag.StringVar(&flagLogLevel, "l", defaultLogLevel, "log level")
	flag.StringVar(&flagDatabaseDSN, "d", "", "db address")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagLogLevel = envLogLevel
	}
	if envDatabaseDSN := os.Getenv("DATABASE_DSN"); envDatabaseDSN != "" {
		flagDatabaseDSN = envDatabaseDSN
	}

	config := &Config{
		Address:     flagRunAddr,
		LogLevel:    flagLogLevel,
		DatabaseDSN: flagDatabaseDSN,
	}

	return config
}

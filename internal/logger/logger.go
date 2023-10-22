// Package logger provides a logger.
package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is available to all codes as a Singleton.
var Log *zap.Logger = zap.NewNop()

// InitLogger initializes logger.
func InitLogger(level string) error {
	// convert text log level to zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()

	cfg.Level = lvl

	cfg.EncoderConfig.EncodeTime = CustomMillisTimeEncoder

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	// set Singleton
	Log = zl

	return nil
}

// CustomMillisTimeEncoder is unix os time encoder.
func CustomMillisTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(`2006-01-02T15:04:05.000207`))
}

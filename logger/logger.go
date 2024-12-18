package logger

import (
	"go.uber.org/zap/zapcore"
	"sync"

	"go.uber.org/zap"
)

var (
	logInstance *zap.Logger
	once        sync.Once
)

// GetLogger initializes and returns a single instance of Zap logger.
func GetLogger() *zap.Logger {
	once.Do(func() {
		config := zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			OutputPaths:      []string{"stdout", "logger/logs/toolbox.log"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:      "time",
				LevelKey:     "level",
				MessageKey:   "msg",
				EncodeLevel:  zapcore.CapitalLevelEncoder,
				EncodeTime:   zapcore.ISO8601TimeEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		logger, err := config.Build()
		if err != nil {
			panic(err)
		}
		logInstance = logger
	})
	return logInstance
}

// SyncLogger flushes any buffered log entries.
func SyncLogger() {
	if logInstance != nil {
		_ = logInstance.Sync()
	}
}

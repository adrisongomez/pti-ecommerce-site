package loggers

import (
	"os"

	"go.uber.org/zap"
)

func CreateLogger(serviceName string) *zap.Logger {
	var logger *zap.Logger
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}
	logger.With(
		zap.String("serviceName", serviceName),
	)
	return logger
}

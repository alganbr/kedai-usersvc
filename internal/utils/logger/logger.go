package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger() Logger {
	logger, _ := zap.NewDevelopment()

	return Logger{
		Zap: logger.Sugar(),
	}
}

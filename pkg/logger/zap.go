package logger

import "go.uber.org/zap"

var logger *Logger

type Logger struct {
	*zap.Logger
}

func Init() {
	logger = &Logger{initZapLogger()}
}

func initZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func Debug(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)

}
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

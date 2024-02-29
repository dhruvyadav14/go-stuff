package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level defines the available log levels.
type Level zap.Level

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// NewLogger creates a new asynchronous structured logger.
func NewLogger(cfg zap.Config) (*zap.Logger, error) {
	return cfg.Build()
}

// NewDevelopmentLogger creates a development logger with debug level enabled.
func NewDevelopmentLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	return NewLogger(cfg)
}

// NewProductionLogger creates a production logger with info level as default.
func NewProductionLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	return NewLogger(cfg)
}

// Field creates a field with the given key and value.
func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// WithFields creates a new logger with the provided fields.
func WithFields(logger *zap.Logger, fields []zap.Field) *zap.Logger {
	return logger.With(fields...)
}

// Debug logs a message at the debug level.
func Debug(l *zap.Logger, msg string, fields ...zap.Field) {
  l.Debug(msg, fields...) // Call the existing Debug method
}

// Info logs a message at the info level.
func (l *zap.Logger) Info(msg string, fields ...zap.Field) {
	l.Infow(msg, fields...)
}

// Warn logs a message at the warn level.
func (l *zap.Logger) Warn(msg string, fields ...zap.Field) {
	l.Warnw(msg, fields...)
}

// Error logs a message at the error level.
func (l *zap.Logger) Error(msg string, fields ...zap.Field) {
	l.Errorw(msg, fields...)
}

// Fatal logs a message at the fatal level and exits the program.
func (l *zap.Logger) Fatal(msg string, fields ...zap.Field) {
	l.Fatalw(msg, fields...)
}

// SetLevel sets the minimum enabled log level for the logger.
func SetLevel(logger *zap.Logger, level Level) error {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	logger.Core().SetLevel(atomicLevel)
	return nil
}

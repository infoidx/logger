package logger

import "errors"

var (
	ErrLoggerInstanceNotInit = errors.New("the logger instance has not been initialized")
)

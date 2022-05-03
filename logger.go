package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

var defaultLogger = newLogger(SetStdOutPut, SetInfoLevel, SetReportCallerRTrue, SetJSONFormatterWithTrace)

func newLogger(configures ...Configure) *logrus.Logger {
	logger := logrus.New()
	for _, configure := range configures {
		configure(logger)
	}
	return logger
}

func NewLogger(configures ...Configure) *logrus.Logger {
	allConfigures := make([]Configure, 0)
	// 设置输出,设置日志级别,设置formatter
	allConfigures = append(configures, SetStdOutPut, SetInfoLevel, SetReportCallerRTrue, SetJSONFormatterWithTrace)
	allConfigures = append(allConfigures, configures...)
	return newLogger(allConfigures...)
}

func WithContext(ctx context.Context) *logrus.Entry {
	return logrus.NewEntry(defaultLogger).WithContext(ctx)
}

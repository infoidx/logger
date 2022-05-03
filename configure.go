package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// 实现类似mesher下的defaulter Logger ,
// 实现WithContext方法完成链路
// 实现调用函数的插件。有现成的插件

type Configure func(logger *logrus.Logger) error

// SetInfoLevel 设置Info日志级别
var SetInfoLevel = SetLevel(logrus.InfoLevel)

// SetStdOutPut 设置输出
var SetStdOutPut = SetOutPut(os.Stdout)

// SetReportCallerRTrue 设置追踪信息
var SetReportCallerRTrue = SetReportCaller(true)

// SetJSONFormatterWithTrace 设置格式化方法
var SetJSONFormatterWithTrace = SetFormatter(NewJSONFormatterWithTrace())

// SetLevel 设置日志级别
var SetLevel = func(v logrus.Level) Configure {
	return func(logger *logrus.Logger) error {
		if logger == nil {
			return ErrLoggerInstanceNotInit
		}
		logger.SetLevel(v)
		return nil
	}
}

// SetOutPut 设置输出设备
var SetOutPut = func(writer io.Writer) Configure {
	return func(logger *logrus.Logger) error {
		if logger == nil {
			return ErrLoggerInstanceNotInit
		}
		logger.SetOutput(writer)
		return nil
	}
}

var SetReportCaller = func(b bool) Configure {
	return func(logger *logrus.Logger) error {
		if logger == nil {
			return ErrLoggerInstanceNotInit
		}
		logger.SetReportCaller(b)
		return nil
	}
}

// SetFormatter 设置格式化方法
var SetFormatter = func(formatter logrus.Formatter) Configure {
	return func(logger *logrus.Logger) error {
		if logger == nil {
			return ErrLoggerInstanceNotInit
		}
		logger.SetFormatter(formatter)
		return nil
	}
}

package logger

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var defaultJsonFormatter = &logrus.JSONFormatter{
	TimestampFormat:  "2006-01-02 15:04:05",
	DisableTimestamp: true,
	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		filename := path.Base(frame.File)
		return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
	},
	//CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
	//	//处理文件名
	//	fileName := path.Base(frame.File)
	//	return frame.Function, fileName
	//},
	PrettyPrint: true,
}

type jSONFormatterConfigure func(*logrus.JSONFormatter)

var setJSONFormatterPrettyPrint = func(b bool) jSONFormatterConfigure {
	return func(formatter *logrus.JSONFormatter) {
		formatter.PrettyPrint = b
	}
}

func NewJSONFormatterWithTrace() *JSONFormatterWithTrace {
	return &JSONFormatterWithTrace{
		JSONFormatter: defaultJsonFormatter,
	}
}

type JSONFormatterWithTrace struct {
	*logrus.JSONFormatter
}

func (f *JSONFormatterWithTrace) Format(entry *logrus.Entry) ([]byte, error) {
	// TODO 增加trace
	// TODO entry.Data = data
	return f.JSONFormatter.Format(entry)
}

package util

import (
	"github.com/sirupsen/logrus"
)

//To consider alternative to this, given how this is in maintance mode and no longer actively extended

type DefaultFieldsFormatter struct {
	WrappedFormatter logrus.Formatter
	DefaultFields    logrus.Fields
	PrintLineNumber  bool
}

func Init(formatter **DefaultFieldsFormatter, component string) {

	if *formatter == nil {
		*formatter = &DefaultFieldsFormatter{
			PrintLineNumber: true,
			DefaultFields:   logrus.Fields{"component": component},
		}
	}

	if (*formatter).WrappedFormatter == nil {
		(*formatter).WrappedFormatter = &logrus.JSONFormatter{}
	}

	logrus.SetFormatter((*formatter))
	logrus.SetReportCaller((*formatter).PrintLineNumber)
}

func (f *DefaultFieldsFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	data := make(logrus.Fields, len(entry.Data))

	return f.WrappedFormatter.Format(&logrus.Entry{
		Logger:  entry.Logger,
		Data:    data,
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
		Caller:  entry.Caller,
	})

}

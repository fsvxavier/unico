package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	rootLogger *logrus.Logger
)

// GenericLogger represents log struct
type GenericLogger struct {
	Log           *logrus.Entry
	Hostname      string
	Module        string
	OperationName string
}

func initLogger() {
	rootLogger = logrus.New()
	rootLogger.SetNoLock()
	rootLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	rootLogger.SetLevel(getLogLevel(LogLevel))
}

// GetLogger return a initialized log
func (g *GenericLogger) GetLogger() {
	if rootLogger == nil {
		initLogger()
	}
	//g.Hostname = Hostname
	g.Log = rootLogger.WithFields(logrus.Fields{
		"environment": Env,
		//	"hostname":    g.Hostname,
		"version": Version,
		"app":     App,
		"squad":   Squad,
		"tribe":   Tribe,
		"module":  g.Module,
	})
}

// LogIt log a new message to stdout
func (g *GenericLogger) LogIt(severity, message string, fields map[string]interface{}) {
	logger := g.Log
	logger = logger.WithFields(logrus.Fields{
		"severity": severity,
		"operation": logrus.Fields{
			"name": g.OperationName,
		},
	})
	if fields != nil {
		logger = logger.WithFields(fields)
	}
	switch severity {
	case "ERROR":
		logger.Error(message)
	case "INFO":
		logger.Info(message)
	case "WARN":
		logger.Warn(message)
	case "FATAL":
		logger.Fatal(message)
	case "DEBUG":
		logger.Debug(message)
	case "PANIC":
		logger.Panic(message)
	default:
		logger.Info(message)
	}
}
func getLogLevel(systemVariable string) logrus.Level {
	switch systemVariable {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	case "PANIC":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

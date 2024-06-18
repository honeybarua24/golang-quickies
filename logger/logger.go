package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

type MyCustomHook struct {
	RequestID string
	UserID    string
	Method    string
}

func (hook *MyCustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *MyCustomHook) Fire(entry *logrus.Entry) error {
	entry.Data["custom_field"] = "custom_value"
	entry.Data["request_id"] = hook.RequestID
	entry.Data["user_id"] = hook.UserID
	entry.Data["method"] = hook.Method
	return nil
}

// InitLogger initializes the logger instance
func InitLogger() {
	Logger = logrus.New()
	// Set log level
	Logger.SetLevel(logrus.DebugLevel)
	// Set output to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.Warn("Failed to log to file, using default stderr")
	}
	// Set the output format to JSON
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// Add hooks (optional, for extended functionality)
	Logger.AddHook(&MyCustomHook{})
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}

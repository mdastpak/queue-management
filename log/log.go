package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance.
var Logger = logrus.New()

// Init initializes the logger with standard settings.
func Init() {
	// Set the output to stdout
	Logger.SetOutput(os.Stdout)
	// Set the log level to Info
	Logger.SetLevel(logrus.InfoLevel)
	// Use JSON formatter for structured logging
	Logger.SetFormatter(&logrus.JSONFormatter{})
}

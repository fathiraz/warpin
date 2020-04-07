package tools

import (
	logs "github.com/sirupsen/logrus"
)

const (
	InfoLevel  = "info"
	DebugLevel = "debug"
	ErrorLevel = "error"
	PanicLevel = "panic"
)

// InitLogs to init our logger
func InitLogs() {
	logs.SetFormatter(&logs.TextFormatter{})
}

// Log function to use in our app
func Log(level string, message string, ctx string, scope string) {
	entry := SetLogMessage(message, ctx, scope)

	switch level {
	case InfoLevel:
		entry.Info()
	case DebugLevel:
		entry.Debug()
	case ErrorLevel:
		entry.Error()
	case PanicLevel:
		entry.Panic()
	}
}

// setLogMessage function to set context log
func SetLogMessage(message string, context string, scope string) *logs.Entry {
	contextLog := logs.WithFields(logs.Fields{
		"message": message,
		"context": context,
		"scope":   scope,
	})

	return contextLog
}

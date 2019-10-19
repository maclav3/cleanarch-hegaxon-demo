package log

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

const (
	// CommandHandlerKey is used by command handlers to annotate logs with their name.
	CommandHandlerKey = "command_handler"
	// CommandKey is used by command handlers to annotate logs with the command passed to the handler.
	CommandKey = "command"
	// QueryHandlerKey is used by query handlers to annotate logs with their name.
	QueryHandlerKey = "query_handler"
	// QueryKey is used by query handlers to annotate logs with the query passed to the handler.
	QueryKey = "query"
)

type Logger interface {
	WithField(key string, value interface{}) Logger
	WithJSON(key string, value interface{}) Logger
	WithError(err error) Logger

	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}

type logrusLogger struct {
	entry *logrus.Entry
}

// NewLogger is an implementation of Logger by sirupsen/logrus.
func NewLogger(name string) Logger {
	return &logrusLogger{
		entry: logrus.NewEntry(logrus.New()).WithField("name", name),
	}
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	newLogger := new(logrusLogger)
	newLogger.entry = l.entry.WithField(key, value)
	return newLogger
}

func (l *logrusLogger) WithJSON(key string, value interface{}) Logger {
	b, err := json.Marshal(value)
	if err != nil {
		return l.WithField(key, "<could not marshal to json>")
	}
	return l.WithField(key, string(b))
}

func (l *logrusLogger) WithError(err error) Logger {
	newLogger := new(logrusLogger)
	newLogger.entry = l.entry.WithError(err)
	return newLogger
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.entry.Info(args)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.entry.Error(args)
}

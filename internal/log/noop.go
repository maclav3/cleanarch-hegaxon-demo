package log

type noopLogger struct{}

func NewNoopLogger() Logger {
	return noopLogger{}
}

func (noopLogger) WithField(key string, value interface{}) Logger {
	return noopLogger{}
}

func (noopLogger) WithError(err error) Logger {
	return noopLogger{}
}

func (noopLogger) Debug(args ...interface{}) {
}

func (noopLogger) Info(args ...interface{}) {
}

func (noopLogger) Error(args ...interface{}) {
}

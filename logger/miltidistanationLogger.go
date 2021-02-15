package logger

import (
	"github.com/Artie18/go-logger/emitter"
	"github.com/Artie18/go-logger/level"
	"github.com/Artie18/go-logger/message"
	"github.com/Artie18/go-logger/message/multifield"
)

type MultiDestinationLogger struct {
	emitters []emitter.Emitter
	constantlyLoggedFields map[string]string
}

func NewMultiDestinationLogger(emitters []emitter.Emitter) *MultiDestinationLogger {
	l := new(MultiDestinationLogger)
	l.emitters = emitters
	l.constantlyLoggedFields = map[string]string{}

	return l
}

func (l *MultiDestinationLogger) Log(level level.LogLevel, msg string) message.Message {
	return multifield.NewMessage(level, msg, l.emitters)
}

func (l *MultiDestinationLogger) Info(msg string) message.Message {
	return multifield.NewMessage(level.LevelInfo, msg, l.emitters)
}

func (l *MultiDestinationLogger) Warning(msg string) message.Message {
	return multifield.NewMessage(level.LevelWarning, msg, l.emitters)
}

func (l *MultiDestinationLogger) Error(msg string) message.Message {
	return multifield.NewMessage(level.LevelError, msg, l.emitters)
}

func (l *MultiDestinationLogger) AddConstantField(key string, value string) {
	l.constantlyLoggedFields[key] = value
}

func (l *MultiDestinationLogger) RemoveConstantField(key string) {
	delete(l.constantlyLoggedFields, key)
}

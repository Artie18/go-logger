package logger

import (
	"logger/level"
	"logger/message"
)

type Logger interface {
	Log(level level.LogLevel, msg string) message.Message
	Info(msg string) message.Message
	Warning(msg string) message.Message
	Error(msg string) message.Message
	AddConstantField(key string, value string)
	RemoveConstantField(key string)
}







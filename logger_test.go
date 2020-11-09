package logger

import (
	"logger/emitter"
	"logger/emitter/console"
	"logger/level"
	"testing"
)

func TestLogger_Write(t *testing.T) {
	emitters := []emitter.Emitter{
		console.NewConsole(),
	}

	l := NewMultiDestinationLogger(emitters)

	l.Log(level.LevelInfo, "Hello World").FieldInt("number", 42).Tag("SUPER DUPER WARNING").Write()
	l.Info("Info World").Write()
	l.Error("Error World").Write()
	l.Warning("Warning world").Write()
}

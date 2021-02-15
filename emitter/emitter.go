package emitter

import (
	"github.com/Artie18/go-logger/message"
)

type Emitter interface {
	Emit(message message.Message) (err error)
}

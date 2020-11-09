package emitter

import (
	"logger/message"
)

type Emitter interface {
	Emit(message message.Message) (err error)
}

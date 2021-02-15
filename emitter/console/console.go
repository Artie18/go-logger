package console

import (
	"fmt"
	"github.com/Artie18/go-logger/message"
)

type Console struct {

}

func NewConsole() *Console {
	return &Console{}
}

func (c Console) Emit(message message.Message) (err error) {
	_, err = fmt.Println(message.String())

	return
}



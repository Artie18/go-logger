package console

import (
	"fmt"
	"logger/message"
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



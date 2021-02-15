package multifield

import (
	"fmt"
	"github.com/Artie18/go-logger/emitter"
	"github.com/Artie18/go-logger/level"
	"github.com/Artie18/go-logger/message"
	"strings"
)

type Message struct {
	fields   Fields
	message  string
	level    level.LogLevel
	emitters []emitter.Emitter
	tags 	 *Tags
}

func NewMessage(level level.LogLevel, msg string, emitters []emitter.Emitter) *Message {
	m := new(Message)
	m.fields = map[string]interface{}{}
	m.level = level
	m.message = msg
	m.emitters = emitters
	m.tags = NewTags()

	return m
}

func (m *Message) FieldInt(key string, field int) message.Message {
	return m.addField(key, field)
}

func (m *Message) FieldString(key string, field string) message.Message {
	return m.addField(key, field)
}

func (m *Message) FieldFloat64(key string, field float64) message.Message {
	return m.addField(key, field)
}

func (m *Message) FieldFloat32(key string, field float32) message.Message {
	return m.addField(key, field)
}

func (m *Message) addField(key string, value interface{}) message.Message {
	m.fields[key] = value

	return m
}

func (m *Message) Fields(mapOfFields map[string]interface{}) message.Message {
	for key, value := range mapOfFields {
		m.fields[key] = value
	}
	return m
}

func (m *Message) Write() (err error) {
	for _, e := range m.emitters {
		e.Emit(m)
	}

	return
}

func (m *Message) Tag(tag string) message.Message {
	m.tags.AppendOne(tag)

	return m
}

func (m *Message) Tags(tags []string) message.Message {
	m.tags.Append(tags)

	return m
}

func (m *Message) String() string {
	var sb strings.Builder

	m.appendLevels(&sb)
	m.appendTags(&sb)
	m.appendFields(&sb)
	m.appendMessage(&sb)

	return sb.String()
}

func (m *Message) appendMessage(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("message=\"%s\"", m.message))
}

func (m *Message) appendFields(sb *strings.Builder) {
	if len(m.fields) > 0 {
		sb.WriteString(fmt.Sprintf("%s ", m.fields.String()))
	}
}

func (m *Message) appendLevels(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%s ", m.level.String()))
}

func (m *Message) appendTags(sb *strings.Builder) {
	if tagsString := m.tags.String(); len(tagsString) != 0 {
		sb.WriteString(fmt.Sprintf("%s ", m.tags.String()))
	}
}




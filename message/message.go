package message

type Message interface {
	Tag(tag string) Message
	Tags(tags []string) Message
	FieldInt(key string, field int) Message
	FieldString(key string, field string) Message
	FieldFloat64(key string, field float64) Message
	FieldFloat32(key string, field float32) Message
	Fields(mapOfFields map[string]interface{}) Message
	Write() (err error)
	String() string
}

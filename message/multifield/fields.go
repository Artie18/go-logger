package multifield

import (
	"fmt"
	"strings"
)

type Fields map[string]interface{}

func (f *Fields) String() string {
	var sb strings.Builder
	for key, value := range *f {
		switch v := value.(type) {
		case string:
			sb.WriteString(fmt.Sprintf("%s=\"%v\"", key, v))
		default:
			sb.WriteString(fmt.Sprintf("%s=%v", key, v))
		}
	}
	return sb.String()

}
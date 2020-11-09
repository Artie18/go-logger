package multifield

import (
	"fmt"
	"strings"
)

type Tags struct {
	tags []string
	longestTag int
}

func NewTags() *Tags {
	t := new(Tags)
	t.tags = []string{}
	t.longestTag = 0
	return t
}

func (t *Tags) AppendOne(tag string) {
	t.tags = append(t.tags, tag)

	if len(tag) > t.longestTag {
		t.longestTag = len(tag)
	}
}

func (t *Tags) Append(tags []string) {
	t.tags = append(t.tags, tags...)

	for _, tag := range t.tags {
		if len(tag) > t.longestTag {
			t.longestTag = len(tag)
		}
	}
}

func (t *Tags) LongestTagSize() int {
	return t.longestTag
}

func (t *Tags) IsEmpty() bool {
	return t.longestTag <= 0
}

// String returns all of the tags in a format [%s] [%s] ...
// If no tags present returns empty string
func (t *Tags) String() string {
	if t.longestTag == 0 {
		return ""
	}
	return fmt.Sprintf("%s%s%s", "[", strings.Join(t.tags, "] ["), "]")
}


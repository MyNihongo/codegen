package codegen

import (
	"fmt"
	"strings"
)

type commentBlock struct {
	value string
}

// CommentF creates a new comment statement according to a format
func (f *file) CommentF(format string, args ...interface{}) *file {
	f.append(commentF(format, args...))
	return f
}

// commentF creates a new comment statement according to a format
func commentF(format string, args ...interface{}) block {
	return comment(fmt.Sprintf(format, args...))
}

// comment creates a new comment statement
func comment(value string) block {
	return &commentBlock{value: value}
}

func (c *commentBlock) write(sb *strings.Builder) {
	writeNewLineF(sb, "// %s", c.value)
}

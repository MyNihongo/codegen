package codegen

import (
	"fmt"
	"strings"
)

type commentBlock struct {
	value string
}

// CommentF creates a new comment statement according to a format
func (f *File) CommentF(format string, args ...interface{}) *File {
	f.append(newCommentF(format, args...))
	return f
}

func newCommentF(format string, args ...interface{}) Block {
	return newComment(fmt.Sprintf(format, args...))
}

func newComment(value string) Block {
	return &commentBlock{value: value}
}

func (c *commentBlock) write(sb *strings.Builder) {
	writeNewLineF(sb, "// %s", c.value)
}

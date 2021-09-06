package codegen

import (
	"fmt"
	"strings"
)

type commentStmt struct {
	value string
}

// CommentF creates a new comment statement according to a format
func (f *file) CommentF(format string, args ...interface{}) {
	f.append(commentF(format, args...))
}

// commentF creates a new comment statement according to a format
func commentF(format string, args ...interface{}) *stmt {
	return comment(fmt.Sprintf(format, args...))
}

// comment creates a new comment statement
func comment(value string) *stmt {
	prt := new(stmt)
	*prt = &commentStmt{value: value}
	return prt
}

func (c *commentStmt) write(sb *strings.Builder) {
	writeNewLineF(sb, "// %s", c.value)
}

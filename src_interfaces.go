package codegen

import "strings"

type block interface {
	write(sb *strings.Builder)
}

type stmt interface {
	writeStmt(sb *strings.Builder)
}

type value interface {
	isPointer() bool
	writeValue(sb *strings.Builder)
}

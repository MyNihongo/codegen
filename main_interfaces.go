package codegen

import "strings"

type block interface {
	write(sb *strings.Builder)
}

type stmt interface {
	// writeStmt writes a new statement. Returns True if a new line should be appended after it. Returns false otherwise.
	writeStmt(sb *strings.Builder) bool
}

type value interface {
	writeValue(sb *strings.Builder)
	isPointer() bool
}

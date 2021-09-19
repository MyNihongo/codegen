package codegen

import "strings"

type Block interface {
	write(sb *strings.Builder)
}

type Stmt interface {
	// writeStmt writes a new statement. Returns True if a new line should be appended after it. Returns false otherwise.
	writeStmt(sb *strings.Builder) bool
}

type Value interface {
	writeValue(sb *strings.Builder)
	isPointer() bool
}

type caller interface {
	getCall() Value
}

package codegen

import "strings"

type block interface {
	write(sb *strings.Builder)
}

type stmt interface {
	writeStmt(sb *strings.Builder)
}

type value interface {
	Field(fieldName string) value

	writeValue(sb *strings.Builder)
}

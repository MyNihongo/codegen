package codegen

import (
	"strings"
)

type newValue struct {
	alias      string
	identifier string
}

func New(identifier string) *newValue {
	return &newValue{
		identifier: identifier,
	}
}

func QualNew(alias, identifier string) *newValue {
	return &newValue{
		alias:      alias,
		identifier: identifier,
	}
}

func (n *newValue) writeValue(sb *strings.Builder) {
	sb.WriteString("new(")

	writeAlias(sb, n.alias)
	sb.WriteString(n.identifier)

	sb.WriteByte(')')
}

func isPointer() bool {
	return false
}

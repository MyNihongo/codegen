package codegen

import (
	"strings"
)

type newValue struct {
	alias      string
	identifier string
}

// New creates a new pointer initialisation value
func New(identifier string) *newValue {
	return &newValue{
		identifier: identifier,
	}
}

// QualNew creates a new pointer initialisation value with an alias
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

func (n *newValue) isPointer() bool {
	return false
}

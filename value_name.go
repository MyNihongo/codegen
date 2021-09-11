package codegen

import "strings"

type nameValue struct {
	alias      string
	identifier string
	isPointer  bool
}

func qualName(alias, identifier string) *nameValue {
	return &nameValue{alias: alias, identifier: identifier}
}

func (n *nameValue) pointer() {
	n.isPointer = true
}

func (n *nameValue) getTypeName() string {
	var sb strings.Builder
	n.writeValue(&sb)

	return sb.String()
}

func (n *nameValue) writeValue(sb *strings.Builder) {
	if n.isPointer {
		sb.WriteByte('*')
	}

	writeAlias(sb, n.alias)
	sb.WriteString(n.identifier)
}

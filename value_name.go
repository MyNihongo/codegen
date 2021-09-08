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

func (q *nameValue) pointer() {
	q.isPointer = true
}

func (q *nameValue) writeValue(sb *strings.Builder) {
	if q.isPointer {
		sb.WriteByte('*')
	}

	writeAlias(sb, q.alias)
	sb.WriteString(q.identifier)
}

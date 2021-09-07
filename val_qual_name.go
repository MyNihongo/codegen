package codegen

import "strings"

type qualNameVal struct {
	alias      string
	identifier string
	isPointer  bool
}

func qualName(alias, identifier string) *qualNameVal {
	return &qualNameVal{alias: alias, identifier: identifier}
}

func (q *qualNameVal) pointer() {
	q.isPointer = true
}

func (q *qualNameVal) writeValue(sb *strings.Builder) {
	if q.isPointer {
		sb.WriteByte('*')
	}

	writeAlias(sb, q.alias)
	sb.WriteString(q.identifier)
}

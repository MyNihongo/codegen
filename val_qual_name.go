package codegen

import "strings"

type qualNameVal struct {
	alias      string
	identifier string
}

func NewQualName(alias, identifier string) *qualNameVal {
	return &qualNameVal{alias: alias, identifier: identifier}
}

func (q *qualNameVal) getValue(sb *strings.Builder) {
	if len(q.alias) != 0 {
		writeF(sb, "%s.", q.alias)
	}

	sb.WriteString(q.identifier)
}

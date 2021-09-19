package codegen

import "strings"

type nameHelper struct {
	alias      string
	identifier string
	isPointer  bool
}

func newNameHelper(alias, identifier string) *nameHelper {
	return &nameHelper{alias: alias, identifier: identifier}
}

func (n *nameHelper) pointer() {
	n.isPointer = true
}

func (n *nameHelper) getTypeName() string {
	var sb strings.Builder
	n.writeValue(&sb)

	return sb.String()
}

func (n *nameHelper) writeValue(sb *strings.Builder) {
	if n.isPointer {
		sb.WriteByte('*')
	}

	writeAlias(sb, n.alias)
	sb.WriteString(n.identifier)
}

package codegen

import "strings"

type nameHelper struct {
	alias      string
	identifier string
	isPointer  bool
}

// Alias returns an alias of the name
func (n *nameHelper) Alias() string {
	return n.alias
}

// Identifier returns an identifier of the name
func (n *nameHelper) Identifier() string {
	return n.identifier
}

func newNameHelper(alias, identifier string) *nameHelper {
	return &nameHelper{alias: alias, identifier: identifier}
}

func (n *nameHelper) pointer(isPointer bool) {
	n.isPointer = isPointer
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

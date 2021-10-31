package codegen

import "strings"

type nameHelper struct {
	alias      string
	identifier string
	isPointer  bool
	isArray    bool
}

func newNameHelper(alias, identifier string) *nameHelper {
	return &nameHelper{alias: alias, identifier: identifier}
}

func (n *nameHelper) setIsPointer(isPointer bool) {
	n.isPointer = isPointer
}

func (n *nameHelper) setIsArray(isArray bool) {
	n.isArray = isArray
}

func (n *nameHelper) getTypeName() string {
	var sb strings.Builder
	n.wr(&sb)

	return sb.String()
}

func (n *nameHelper) isValid() bool {
	return len(n.identifier) > 0
}

func (n *nameHelper) String() string {
	var sb strings.Builder
	n.wr(&sb)

	return sb.String()
}

func (n *nameHelper) wr(sb *strings.Builder) {
	if n.isArray {
		sb.WriteString("[]")
	}

	if n.isPointer {
		sb.WriteByte('*')
	}

	writeAlias(sb, n.alias)
	sb.WriteString(n.identifier)
}

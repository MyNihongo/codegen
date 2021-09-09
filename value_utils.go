package codegen

import "strings"

func writeValues(sb *strings.Builder, vals []Value) {
	for i, v := range vals {
		if i != 0 {
			sb.WriteByte(',')
		}

		v.writeValue(sb)
	}
}

func writePointerValueAccess(sb *strings.Builder, val Value) {
	isPointer := val.isPointer()
	if isPointer {
		sb.WriteByte('(')
	}

	val.writeValue(sb)

	if isPointer {
		sb.WriteByte(')')
	}
}

func writeFuncCall(sb *strings.Builder, name string, vals []Value) {
	writeF(sb, "%s(", name)
	writeValues(sb, vals)
	sb.WriteByte(')')
}

func writeAlias(sb *strings.Builder, alias string) {
	if len(alias) != 0 {
		writeF(sb, "%s.", alias)
	}
}

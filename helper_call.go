package codegen

import "strings"

type callHelper struct {
	name string
	args []Value
}

func newCallHelper(name string, args []Value) *callHelper {
	return &callHelper{
		name: name,
		args: args,
	}
}

func (c *callHelper) wr(sb *strings.Builder) {
	writeF(sb, "%s(", c.name)
	writeValues(sb, c.args)
	sb.WriteByte(')')
}

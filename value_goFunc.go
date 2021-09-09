package codegen

import "strings"

type goFuncValue struct {
	name string
	args []value
}

// Len creates a new function call of the Go build-in function `len()`
func Len(val value) *goFuncValue {
	return newGoFunc("len", val)
}

// Equals compares a value of the go function for equality
func (g *goFuncValue) Equals(val value) *equalsValue {
	return newEquals(g, val, true)
}

// Equals compares a value of the go function for not being equal
func (g *goFuncValue) NotEquals(val value) *equalsValue {
	return newEquals(g, val, false)
}

func newGoFunc(name string, args ...value) *goFuncValue {
	return &goFuncValue{
		name: name,
		args: args,
	}
}

func (g *goFuncValue) writeValue(sb *strings.Builder) {
	writeF(sb, "%s(", g.name)
	writeValues(sb, g.args)
	sb.WriteByte(')')
}

func (g *goFuncValue) isPointer() bool {
	return false
}

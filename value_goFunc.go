package codegen

import "strings"

type goFuncValue struct {
	name string
	args []Value
}

// Len creates a new function call of the Go build-in function `len()`
func Len(val Value) *goFuncValue {
	return newGoFunc("len", val)
}

// Equals compares a value of the go function for equality
func (g *goFuncValue) Equals(val Value) *equalsValue {
	return newEquals(g, val, true)
}

// Equals compares a value of the go function for not being equal
func (g *goFuncValue) NotEquals(val Value) *equalsValue {
	return newEquals(g, val, false)
}

func newGoFunc(name string, args ...Value) *goFuncValue {
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

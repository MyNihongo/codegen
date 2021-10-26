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

// Make creates a new function call of the Go build-in function `make()` for an empty slice
func Make(sliceType *TypeDecl) *goFuncValue {
	return MakeWithCount(sliceType, 0)
}

// Make creates a new function call of the Go build-in function `make()` for a slice with count
func MakeWithCount(sliceType *TypeDecl, count int) *goFuncValue {
	sliceType.Array()
	typeString := Identifier(sliceType.name.String())
	return newGoFunc("make", typeString, Int(count))
}

// Equals compares a value of the go function for equality
func (g *goFuncValue) Equals(val Value) *comparisonValue {
	return newEquals(g, val, cmpType_Equals)
}

// Equals compares a value of the go function for not being equal
func (g *goFuncValue) NotEquals(val Value) *comparisonValue {
	return newEquals(g, val, cmpType_NotEquals)
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

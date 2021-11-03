package codegen

import "strings"

type goFuncValue struct {
	name string
	args []Value
}

// Len creates a new function call of the Go built-in function `len()`
func Len(val Value) *goFuncValue {
	return newGoFunc("len", val)
}

// MakeSlice creates a new function call of the Go built-in function `make()` for an empty slice
func MakeSlice(sliceType *TypeDecl) *goFuncValue {
	return MakeSliceWithCount(sliceType, 0)
}

// Make creates a new function call of the Go built-in function `make()` for a slice with count
func MakeSliceWithCount(sliceType *TypeDecl, count int) *goFuncValue {
	sliceType.Array()
	typeString := Identifier(sliceType.name.String())
	return newGoFunc("make", typeString, Int(count))
}

// Append creates a new function call of the built-in function `append`
func Append(sliceValue Value, elementValues ...Value) *goFuncValue {
	vals := make([]Value, len(elementValues)+1)
	vals[0] = sliceValue

	for i, v := range elementValues {
		vals[i+1] = v
	}

	return newGoFunc("append", vals...)
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

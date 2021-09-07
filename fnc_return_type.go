package codegen

import "strings"

type returnType struct {
	name      *qualNameVal
	isPointer bool
}

func NewReturnType(name string) *returnType {
	return &returnType{name: NewQualName("", name)}
}

func NewQualReturnType(alias, name string) *returnType {
	return &returnType{name: NewQualName(alias, name)}
}

func NewReturnTypeError() *returnType {
	return NewReturnType("error")
}

func (r *returnType) Pointer() *returnType {
	r.isPointer = true
	return r
}

func (r *returnType) wr(sb *strings.Builder) {
	if r.isPointer {
		sb.WriteByte('*')
	}

	r.name.writeValue(sb)
}

func writeReturnTypes(sb *strings.Builder, returnTypes []*returnType) {
	if count := len(returnTypes); count == 0 {
		return
	} else if count == 1 {
		returnTypes[0].wr(sb)
	} else {
		sb.WriteByte('(')

		for i, r := range returnTypes {
			if i != 0 {
				sb.WriteByte(',')
			}

			r.wr(sb)
		}

		sb.WriteByte(')')
	}
}

package codegen

import "strings"

type returnType struct {
	name *qualNameVal
}

func ReturnType(name string) *returnType {
	return &returnType{name: NewQualName("", name)}
}

func QualReturnType(alias, name string) *returnType {
	return &returnType{name: NewQualName(alias, name)}
}

func ReturnTypeError() *returnType {
	return ReturnType("error")
}

func (r *returnType) Pointer() *returnType {
	r.name.pointer()
	return r
}

func (r *returnType) wr(sb *strings.Builder) {
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

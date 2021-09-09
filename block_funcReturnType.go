package codegen

import "strings"

type returnType struct {
	name *nameValue
}

// ReturnType creates a new return type for a function
func ReturnType(name string) *returnType {
	return &returnType{name: qualName("", name)}
}

// QualReturnType creates a new return type with an alias of an imported package
func QualReturnType(alias, name string) *returnType {
	return &returnType{name: qualName(alias, name)}
}

// ReturnTypeError create a new return type of type `error`
func ReturnTypeError() *returnType {
	return ReturnType("error")
}

// Pointer turns the return type into a pointer value
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

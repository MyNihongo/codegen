package codegen

import "strings"

type ReturnTypeDecl struct {
	*nameHelper
}

// ReturnType creates a new return type for a function
func ReturnType(name string) *ReturnTypeDecl {
	return &ReturnTypeDecl{nameHelper: newNameHelper("", name)}
}

// QualReturnType creates a new return type with an alias of an imported package
func QualReturnType(alias, name string) *ReturnTypeDecl {
	return &ReturnTypeDecl{nameHelper: newNameHelper(alias, name)}
}

// ReturnTypeError create a new return type of type `error`
func ReturnTypeError() *ReturnTypeDecl {
	return ReturnType("error")
}

// Pointer turns the return type into a pointer value
func (r *ReturnTypeDecl) Pointer() *ReturnTypeDecl {
	r.SetIsPointer(true)
	return r
}

// SetIsPointer sets whether or not a return type is a pointer
func (r *ReturnTypeDecl) SetIsPointer(isPointer bool) *ReturnTypeDecl {
	r.nameHelper.pointer(isPointer)
	return r
}

func (r *ReturnTypeDecl) wr(sb *strings.Builder) {
	r.nameHelper.writeValue(sb)
}

func writeReturnTypes(sb *strings.Builder, returnTypes []*ReturnTypeDecl) {
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

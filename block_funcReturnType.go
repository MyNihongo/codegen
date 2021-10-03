package codegen

import "strings"

type ReturnTypeDecl struct {
	name *nameHelper
}

// TypeName gets a type name of the return declaration
func (r *ReturnTypeDecl) TypeName() string {
	return r.name.identifier
}

// TypeAlias gets a type alias (if any) of the return declaration
func (r *ReturnTypeDecl) TypeAlias() string {
	return r.name.alias
}

// ReturnType creates a new return type for a function
func ReturnType(name string) *ReturnTypeDecl {
	return &ReturnTypeDecl{name: newNameHelper("", name)}
}

// QualReturnType creates a new return type with an alias of an imported package
func QualReturnType(alias, name string) *ReturnTypeDecl {
	return &ReturnTypeDecl{name: newNameHelper(alias, name)}
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
	r.name.pointer(isPointer)
	return r
}

func (r *ReturnTypeDecl) wr(sb *strings.Builder) {
	r.name.writeValue(sb)
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

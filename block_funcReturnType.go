package codegen

import "strings"

type ReturnTypeDecl struct {
	name *nameHelper
}

// GetTypeName gets a type name of the return declaration
func (r *ReturnTypeDecl) GetTypeName() string {
	return r.name.identifier
}

// GetTypeAlias gets a type alias (if any) of the return declaration
func (r *ReturnTypeDecl) GetTypeAlias() string {
	return r.name.alias
}

// GetIsPointer gets a flag whether or not the return type is a pointer
func (r *ReturnTypeDecl) GetIsPointer() bool {
	return r.name.isPointer
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
	r.name.setIsPointer(isPointer)
	return r
}

func (r *ReturnTypeDecl) wr(sb *strings.Builder) {
	r.name.writeValue(sb)
}

func (r *ReturnTypeDecl) isValid() bool {
	return r.name.isValid()
}

func writeReturnTypes(sb *strings.Builder, returnTypes []*ReturnTypeDecl) {
	if count := len(returnTypes); count == 0 {
		return
	} else if count == 1 {
		if returnTypes[0].isValid() {
			returnTypes[0].wr(sb)
		}
	} else {
		validCounter := 0
		for _, r := range returnTypes {
			if !r.isValid() {
				continue
			}

			if validCounter != 0 {
				sb.WriteByte(',')
			} else {
				sb.WriteByte('(')
			}

			r.wr(sb)
			validCounter++
		}

		if validCounter != 0 {
			sb.WriteByte(')')
		}
	}
}

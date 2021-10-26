package codegen

import "strings"

type TypeDecl struct {
	name *nameHelper
}

// GetTypeName gets a type name of the return declaration
func (r *TypeDecl) GetTypeName() string {
	return r.name.identifier
}

// GetTypeAlias gets a type alias (if any) of the return declaration
func (r *TypeDecl) GetTypeAlias() string {
	return r.name.alias
}

// GetIsPointer gets a flag whether or not the return type is a pointer
func (r *TypeDecl) GetIsPointer() bool {
	return r.name.isPointer
}

// ReturnType creates a new return type for a function
func ReturnType(name string) *TypeDecl {
	return &TypeDecl{name: newNameHelper("", name)}
}

// QualReturnType creates a new return type with an alias of an imported package
func QualReturnType(alias, name string) *TypeDecl {
	return &TypeDecl{name: newNameHelper(alias, name)}
}

// ReturnTypeError create a new return type of type `error`
func ReturnTypeError() *TypeDecl {
	return ReturnType("error")
}

// Pointer turns the return type into a pointer value
func (r *TypeDecl) Pointer() *TypeDecl {
	r.SetIsPointer(true)
	return r
}

// SetIsPointer sets whether or not a return type is a pointer
func (r *TypeDecl) SetIsPointer(isPointer bool) *TypeDecl {
	r.name.setIsPointer(isPointer)
	return r
}

func (r *TypeDecl) wr(sb *strings.Builder) {
	r.name.writeValue(sb)
}

func (r *TypeDecl) isValid() bool {
	return r.name.isValid()
}

func writeReturnTypes(sb *strings.Builder, returnTypes []*TypeDecl) {
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

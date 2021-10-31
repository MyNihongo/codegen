package codegen

import "strings"

type TypeDecl struct {
	name *nameHelper
}

// Type creates a new type for a function
func Type(name string) *TypeDecl {
	return &TypeDecl{name: newNameHelper("", name)}
}

// QualType creates a new type with an alias of an imported package
func QualType(alias, name string) *TypeDecl {
	return &TypeDecl{name: newNameHelper(alias, name)}
}

// ReturnTypeError create a new type of type `error`
func ReturnTypeError() *TypeDecl {
	return Type("error")
}

// GetTypeName gets a type name of the return declaration
func (t *TypeDecl) GetTypeName() string {
	return t.name.identifier
}

// GetTypeAlias gets a type alias (if any) of the return declaration
func (t *TypeDecl) GetTypeAlias() string {
	return t.name.alias
}

// GetIsPointer gets a flag whether or not the type is a pointer
func (t *TypeDecl) GetIsPointer() bool {
	return t.name.isPointer
}

// Pointer turns the type into a pointer value
func (t *TypeDecl) Pointer() *TypeDecl {
	t.SetIsPointer(true)
	return t
}

// SetIsPointer sets whether or not the type is a pointer
func (t *TypeDecl) SetIsPointer(isPointer bool) *TypeDecl {
	t.name.setIsPointer(isPointer)
	return t
}

// Array converts the type to the array type
func (t *TypeDecl) Array() *TypeDecl {
	return t.SetIsArray(true)
}

// SetIsArray sets whether or not the type is an array
func (t *TypeDecl) SetIsArray(isArray bool) *TypeDecl {
	t.name.setIsArray(isArray)
	return t
}

func (r *TypeDecl) wr(sb *strings.Builder) {
	r.name.wr(sb)
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

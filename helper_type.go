package codegen

import "strings"

type TypeDecl struct {
	name *nameHelper
}

// Type creates a new return type for a function
func Type(name string) *TypeDecl {
	return &TypeDecl{name: newNameHelper("", name)}
}

// QualType creates a new return type with an alias of an imported package
func QualType(alias, name string) *TypeDecl {
	return &TypeDecl{name: newNameHelper(alias, name)}
}

// ReturnTypeError create a new return type of type `error`
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

// GetIsPointer gets a flag whether or not the return type is a pointer
func (t *TypeDecl) GetIsPointer() bool {
	return t.name.isPointer
}

// Pointer turns the return type into a pointer value
func (t *TypeDecl) Pointer() *TypeDecl {
	t.SetIsPointer(true)
	return t
}

// SetIsPointer sets whether or not a return type is a pointer
func (t *TypeDecl) SetIsPointer(isPointer bool) *TypeDecl {
	t.name.setIsPointer(isPointer)
	return t
}

// Array converts the type to the array type
func (t *TypeDecl) Array() *TypeDecl {
	t.name.setIsArray(true)
	return t
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

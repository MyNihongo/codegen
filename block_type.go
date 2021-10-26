package codegen

import (
	"strings"
)

type typeBlock struct {
	name     string
	baseType *nameHelper
}

// Type creates a new type block
func (f *File) Type(typeName, baseType string) *typeBlock {
	t := newType(typeName, "", baseType)
	f.append(t)

	return t
}

// Type creates a new type block with a package alias
func (f *File) QualType(typeName, alias, baseType string) *typeBlock {
	t := newType(typeName, alias, baseType)
	f.append(t)

	return t
}

// Pointer turns the base type to the poiter type
func (t *typeBlock) Pointer() *typeBlock {
	t.SetIsPointer(true)
	return t
}

// SetIsPointer sets whether or not a type is a pointer
func (t *typeBlock) SetIsPointer(isPointer bool) *typeBlock {
	t.baseType.setIsPointer(isPointer)
	return t
}

func newType(name, alias, baseType string) *typeBlock {
	return &typeBlock{
		name:     name,
		baseType: newNameHelper(alias, baseType),
	}
}

func (t *typeBlock) write(sb *strings.Builder) {
	writeF(sb, "type %s ", t.name)
	t.baseType.writeValue(sb)
	newLine(sb)
}

package codegen

import "strings"

type PropertyDecl struct {
	name     string
	typeName *nameHelper
}

// Property creates a new property declaration
func Property(propertyName, typeName string) *PropertyDecl {
	return &PropertyDecl{
		name:     propertyName,
		typeName: newNameHelper("", typeName),
	}
}

// QualProperty creates a new property declaration with a package alias
func QualProperty(propertyName, alias, typeName string) *PropertyDecl {
	return &PropertyDecl{
		name:     propertyName,
		typeName: newNameHelper(alias, typeName),
	}
}

// EmbeddedProp crates a new embedded property
func EmbeddedProperty(typeName string) *PropertyDecl {
	return Property("", typeName)
}

// QualEmbeddedProperty creates a new embedded property with a package alias
func QualEmbeddedProperty(alias, typeName string) *PropertyDecl {
	return QualProperty("", alias, typeName)
}

// Pointer turns the property into a pointer type
func (p *PropertyDecl) Pointer() *PropertyDecl {
	p.typeName.pointer()
	return p
}

func (p *PropertyDecl) wr(sb *strings.Builder) {
	if len(p.name) != 0 {
		writeF(sb, "%s ", p.name)
	}

	p.typeName.writeValue(sb)
}

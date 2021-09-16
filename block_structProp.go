package codegen

import "strings"

type PropertyDecl struct {
	name     string
	typeName *nameValue
}

// Property creates a new property declaration
func Property(propertyName, typeName string) *PropertyDecl {
	return &PropertyDecl{
		name:     propertyName,
		typeName: qualName("", typeName),
	}
}

// QualProperty creates a new property declaration with a package alias
func QualProperty(propertyName, alias, typeName string) *PropertyDecl {
	return &PropertyDecl{
		name:     propertyName,
		typeName: qualName(alias, typeName),
	}
}

// Pointer turns the property into a pointer type
func (p *PropertyDecl) Pointer() *PropertyDecl {
	p.typeName.pointer()
	return p
}

func (p *PropertyDecl) wr(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)
	p.typeName.writeValue(sb)
}

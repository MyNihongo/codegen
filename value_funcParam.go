package codegen

import "strings"

type paramValue struct {
	name     string
	typeName *nameValue
}

// Param creates a new function parameter
func Param(name, typeName string) *paramValue {
	return &paramValue{name: name, typeName: qualName("", typeName)}
}

// QualParam creates a new function parameter with a package alias
func QualParam(name, alias, typeName string) *paramValue {
	return &paramValue{name: name, typeName: qualName(alias, typeName)}
}

// Pointer turns the parameter into a pointer type
func (p *paramValue) Pointer() *paramValue {
	p.typeName.pointer()
	return p
}

func (p *paramValue) writeValue(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)
	p.typeName.writeValue(sb)
}

func writeParams(sb *strings.Builder, params []*paramValue) {
	sb.WriteByte('(')

	for i, p := range params {
		if i != 0 {
			sb.WriteByte(',')
		}

		p.writeValue(sb)
	}

	sb.WriteByte(')')
}

package codegen

import "strings"

type ParamDecl struct {
	name     string
	typeName *nameHelper
}

// Name gets the name of the parameter
func (p *ParamDecl) Name() string {
	return p.name
}

// TypeName gets the type name of the parameter
func (p *ParamDecl) TypeName() string {
	var sb strings.Builder
	p.typeName.writeValue(&sb)

	return sb.String()
}

// Param creates a new function parameter
func Param(name, typeName string) *ParamDecl {
	return &ParamDecl{name: name, typeName: qualName("", typeName)}
}

// QualParam creates a new function parameter with a package alias
func QualParam(name, alias, typeName string) *ParamDecl {
	return &ParamDecl{name: name, typeName: qualName(alias, typeName)}
}

// Pointer turns the parameter into a pointer type
func (p *ParamDecl) Pointer() *ParamDecl {
	p.typeName.pointer()
	return p
}

func (p *ParamDecl) wr(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)
	p.typeName.writeValue(sb)
}

func writeParams(sb *strings.Builder, params []*ParamDecl) {
	sb.WriteByte('(')

	for i, p := range params {
		if i != 0 {
			sb.WriteByte(',')
		}

		p.wr(sb)
	}

	sb.WriteByte(')')
}

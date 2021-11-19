package codegen

import "strings"

type ParamDecl struct {
	name     string
	typeName *nameHelper
}

// Param creates a new function parameter
func Param(name, typeName string) *ParamDecl {
	return &ParamDecl{name: name, typeName: newNameHelper("", typeName)}
}

// QualParam creates a new function parameter with a package alias
func QualParam(name, alias, typeName string) *ParamDecl {
	return &ParamDecl{name: name, typeName: newNameHelper(alias, typeName)}
}

// GetName gets the name of the parameter
func (p *ParamDecl) GetName() string {
	return p.name
}

// GetTypeName gets the type name of the parameter
func (p *ParamDecl) GetTypeName() string {
	return p.typeName.identifier
}

// GetTypeAlias return an alias of the type package (if any) from there the type is imported
func (p *ParamDecl) GetTypeAlias() string {
	return p.typeName.alias
}

// GetIsPointer gets a flag whether or not the parameter type is a pointer
func (p *ParamDecl) GetIsPointer() bool {
	return p.typeName.isPointer
}

// GetIsArray gets a flag whether or not the parameter is an array or not
func (p *ParamDecl) GetIsArray() bool {
	return p.typeName.isArray
}

// GetFullType gets the full string representation of the type
func (p *ParamDecl) GetFullType() string {
	var sb strings.Builder
	p.typeName.wr(&sb)

	return sb.String()
}

// Pointer turns the parameter into a pointer type
func (p *ParamDecl) Pointer() *ParamDecl {
	p.SetIsPointer(true)
	return p
}

// SetIsPointer sets whether or not the parameter is a pointer
func (p *ParamDecl) SetIsPointer(isPointer bool) *ParamDecl {
	p.typeName.setIsPointer(isPointer)
	return p
}

// Array sets the parameter to an array type
func (p *ParamDecl) Array() *ParamDecl {
	return p.SetIsArray(true)
}

// SetIsArray sets whether or not the parameter is an array
func (p *ParamDecl) SetIsArray(isArray bool) *ParamDecl {
	p.typeName.setIsArray(isArray)
	return p
}

func (p *ParamDecl) wr(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)
	p.typeName.wr(sb)
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

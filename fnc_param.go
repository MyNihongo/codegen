package codegen

import "strings"

type paramVal struct {
	name     string
	typeName *qualNameVal
}

func Param(name, typeName string) *paramVal {
	return &paramVal{name: name, typeName: NewQualName("", typeName)}
}

func QualParam(name, alias, typeName string) *paramVal {
	return &paramVal{name: name, typeName: NewQualName(alias, typeName)}
}

func (p *paramVal) Pointer() *paramVal {
	p.typeName.pointer()
	return p
}

func (p *paramVal) writeValue(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)
	p.typeName.writeValue(sb)
}

func writeParams(sb *strings.Builder, params []*paramVal) {
	sb.WriteByte('(')

	for i, p := range params {
		if i != 0 {
			sb.WriteByte(',')
		}

		p.writeValue(sb)
	}

	sb.WriteByte(')')
}

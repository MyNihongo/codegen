package codegen

import "strings"

type paramVal struct {
	name      string
	typeName  *qualNameVal
	isPointer bool
}

func NewParam(name, typeName string) *paramVal {
	return &paramVal{name: name, typeName: NewQualName("", typeName)}
}

func NewQualParam(name string, qualName *qualNameVal) *paramVal {
	return &paramVal{name: name, typeName: qualName}
}

func (p *paramVal) Pointer() *paramVal {
	p.isPointer = true
	return p
}

func (p *paramVal) writeValue(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)

	if p.isPointer {
		sb.WriteByte('*')
	}

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

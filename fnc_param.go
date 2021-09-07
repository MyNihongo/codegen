package codegen

import "strings"

type paramStmt struct {
	name      string
	typeName  *qualNameVal
	isPointer bool
}

func NewParam(name, typeName string) *paramStmt {
	return &paramStmt{name: name, typeName: NewQualName("", typeName)}
}

func NewQualParam(name string, qualName *qualNameVal) *paramStmt {
	return &paramStmt{name: name, typeName: qualName}
}

func (p *paramStmt) Pointer() *paramStmt {
	p.isPointer = true
	return p
}

func (p *paramStmt) writeValue(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)

	if p.isPointer {
		sb.WriteByte('*')
	}

	p.typeName.writeValue(sb)
}

func writeParams(sb *strings.Builder, params []*paramStmt) {
	sb.WriteByte('(')

	for i, p := range params {
		if i != 0 {
			sb.WriteByte(',')
		}

		p.writeValue(sb)
	}

	sb.WriteByte(')')
}

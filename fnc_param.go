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

func (p *paramStmt) getValue(sb *strings.Builder) {
	writeF(sb, "%s ", p.name)

	if p.isPointer {
		sb.WriteByte('*')
	}

	p.typeName.getValue(sb)
}

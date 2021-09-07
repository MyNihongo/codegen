package codegen

import "strings"

type methodBlock struct {
	this *thisVal
	*funcBlock
}

// NewMethod creates a new method block
func (f *file) NewMethod(this *thisVal, name string) *methodBlock {
	method := newMethod(this, name)

	ptr := new(block)
	*ptr = method
	f.append(ptr)

	return method
}

// Params appends method parameters
func (m *methodBlock) Params(params ...*paramVal) *methodBlock {
	m.funcBlock.Params(params...)
	return m
}

// ReturnTypes appends method return parameters
func (m *methodBlock) ReturnTypes(returnTypes ...*returnType) *methodBlock {
	m.funcBlock.ReturnTypes(returnTypes...)
	return m
}

func newMethod(this *thisVal, name string) *methodBlock {
	return &methodBlock{
		this:      this,
		funcBlock: newFunc(name),
	}
}

func (m *methodBlock) write(sb *strings.Builder) {
	sb.WriteString("func (")
	m.this.writeValue(sb)
	writeF(sb, ") %s", m.name)

	writeParams(sb, m.params)
	writeReturnTypes(sb, m.retTypes)
	writeStmts(sb, m.stmts, true)
}

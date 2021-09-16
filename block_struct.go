package codegen

import (
	"strings"
)

type structBlock struct {
	name  string
	props []*PropertyDecl
}

// Struct creates a new struct block
func (f *File) Struct(structName string) *structBlock {
	str := newStruct(structName)
	f.append(str)

	return str
}

// Props adds property declarations to the struct block
func (s *structBlock) Props(properties ...*PropertyDecl) *structBlock {
	s.props = properties
	return s
}

func newStruct(name string) *structBlock {
	return &structBlock{
		name:  name,
		props: make([]*PropertyDecl, 0),
	}
}

func (s *structBlock) write(sb *strings.Builder) {
	writeNewLineF(sb, "type %s struct {", s.name)

	for _, prop := range s.props {
		prop.wr(sb)
		newLine(sb)
	}
	writeByteNewLine(sb, '}')
}

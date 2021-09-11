package codegen

import "strings"

type structValue struct {
	name      string
	props     []*StructPropertyValue
	isAddress bool
}

type StructPropertyValue struct {
	name string
	val  Value
}

// InitStruct creates a new struct initialisation
func InitStruct(structName string) *structValue {
	return &structValue{
		name:  structName,
		props: make([]*StructPropertyValue, 0),
	}
}

// Props adds properties with values to the struct initialisation
func (s *structValue) Props(properties ...*StructPropertyValue) *structValue {
	s.props = properties
	return s
}

// Address returns a pointer to the initialised struct
func (s *structValue) Address() *structValue {
	s.isAddress = true
	return s
}

// PropValue creates a new property with its value
func PropValue(propertyName string, propertyValue Value) *StructPropertyValue {
	return &StructPropertyValue{
		name: propertyName,
		val:  propertyValue,
	}
}

func (s *structValue) writeValue(sb *strings.Builder) {
	if s.isAddress {
		sb.WriteByte('&')
	}

	sb.WriteString(s.name)
	sb.WriteByte('{')

	for i, prop := range s.props {
		if i != 0 {
			sb.WriteByte(',')
		}

		sb.WriteString(prop.name)
		sb.WriteByte(':')
		prop.val.writeValue(sb)
	}

	sb.WriteByte('}')
}

func (s *structValue) isPointer() bool {
	return false
}

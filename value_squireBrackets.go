package codegen

import "strings"

type squireBracketsValue struct {
	val           Value
	valInBrackets Value
}

func newSquireBrackets(val Value, valueInBrackets Value) *squireBracketsValue {
	return &squireBracketsValue{
		val:           val,
		valInBrackets: valueInBrackets,
	}
}

func (s *squireBracketsValue) writeValue(sb *strings.Builder) {
	writePointerValueAccess(sb, s.val)

	sb.WriteByte('[')
	s.valInBrackets.writeValue(sb)
	sb.WriteByte(']')
}

func (s *squireBracketsValue) isPointer() bool {
	return false
}

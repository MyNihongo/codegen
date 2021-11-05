package codegen

import "strings"

type squireBracketsValue struct {
	val           Value
	valInBrackets Value
}

// Assign assigns a value to the squire bracket notation
func (s *squireBracketsValue) Assign(value Value) *assignStmt {
	return newAssignment(s, value)
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

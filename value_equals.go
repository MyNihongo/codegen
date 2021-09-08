package codegen

import "strings"

type equalsValue struct {
	equals   bool
	valRight value
	valLeft  value
}

func newEquals(valRight, valLeft value, equals bool) *equalsValue {
	return &equalsValue{
		equals:   equals,
		valRight: valRight,
		valLeft:  valLeft,
	}
}

func (e *equalsValue) writeValue(sb *strings.Builder) {
	e.valRight.writeValue(sb)

	if e.equals {
		sb.WriteString("==")
	} else {
		sb.WriteString("!=")
	}

	e.valLeft.writeValue(sb)
}

func (e *equalsValue) isPointer() bool {
	return false
}

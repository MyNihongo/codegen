package codegen

import "strings"

type cmpType byte

const (
	cmpType_Equals    cmpType = 1
	cmpType_NotEquals cmpType = 2
	cmpType_LessThan  cmpType = 3
)

type comparisonValue struct {
	cmpType  cmpType
	valRight Value
	valLeft  Value
}

func newEquals(valRight, valLeft Value, cmpType cmpType) *comparisonValue {
	return &comparisonValue{
		cmpType:  cmpType,
		valRight: valRight,
		valLeft:  valLeft,
	}
}

func (e *comparisonValue) writeValue(sb *strings.Builder) {
	e.valRight.writeValue(sb)

	switch e.cmpType {
	case cmpType_Equals:
		sb.WriteString("==")
	case cmpType_NotEquals:
		sb.WriteString("!=")
	case cmpType_LessThan:
		sb.WriteByte('<')
	default:
		panic("unknown comparison type")
	}

	e.valLeft.writeValue(sb)
}

func (e *comparisonValue) isPointer() bool {
	return false
}

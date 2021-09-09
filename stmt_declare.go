package codegen

import "strings"

type declarationValues struct {
	vars    []string
	declare bool
}

type declarationStmt struct {
	varVal *declarationValues
	values []Value
}

// Declare creates a new declaration statement without variable values (will not compile).
// In order to declare variables call `Values()`
func Declare(vars ...string) *declarationValues {
	if len(vars) == 0 {
		panic("no variables are passed for declaration")
	}

	return &declarationValues{
		vars:    vars,
		declare: true,
	}
}

// Completes the declaration block with values
func (v *declarationValues) Values(values ...Value) *declarationStmt {
	if len(values) == 0 {
		panic("no values are provided for variable declaration")
	}

	return &declarationStmt{
		varVal: v,
		values: values,
	}
}

func (d *declarationStmt) writeStmt(sb *strings.Builder) bool {
	for i, v := range d.varVal.vars {
		if i != 0 {
			sb.WriteByte(',')
		}

		sb.WriteString(v)
	}

	if d.varVal.declare {
		sb.WriteString(":=")
	} else {
		sb.WriteByte('=')
	}

	writeValues(sb, d.values)
	return true
}

package codegen

import "strings"

type varsDeclarationStmt struct {
	vars []*VarValue
}

type VarValue struct {
	name     string
	typeName *nameValue
}

// DeclareVars creates a new variable declaration statement
func DeclareVars(vars ...*VarValue) Stmt {
	return &varsDeclarationStmt{
		vars: vars,
	}
}

// Var creates a new variable with a type name
func Var(varName, typeName string) *VarValue {
	return &VarValue{
		name:     varName,
		typeName: qualName("", typeName),
	}
}

// QualVar creates a new variable with a type name and its alias
func QualVar(varName, typeAlias, typeName string) *VarValue {
	return &VarValue{
		name:     varName,
		typeName: qualName(typeAlias, typeName),
	}
}

// Pointer turns the variable type into a pointer
func (v *VarValue) Pointer() *VarValue {
	v.typeName.pointer()
	return v
}

func (v *varsDeclarationStmt) writeStmt(sb *strings.Builder) bool {
	if len(v.vars) != 0 {
		typeMap := make(map[string][]string)

		for _, val := range v.vars {
			typeName := val.typeName.getTypeName()

			if grouping, ok := typeMap[typeName]; ok {
				typeMap[typeName] = append(grouping, val.name)
			} else {
				typeMap[typeName] = []string{val.name}
			}
		}

		for typeName, vars := range typeMap {
			sb.WriteString("var ")

			for i, varName := range vars {
				if i != 0 {
					sb.WriteByte(',')
				}

				sb.WriteString(varName)
			}

			writeNewLineF(sb, " %s", typeName)
		}
	}

	return false
}

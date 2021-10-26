package codegen

import "strings"

type varsDeclarationStmt struct {
	vars []*VarValue
}

type VarValue struct {
	name     string
	typeName *nameHelper
}

// DeclareVars creates a new variable declaration block
func (f *File) DeclareVars(vars ...*VarValue) Block {
	decl := newVars(vars)
	f.append(decl)

	return decl
}

// DeclareVars creates a new variable declaration statement
func DeclareVars(vars ...*VarValue) Stmt {
	return newVars(vars)
}

// Var creates a new variable with a type name
func Var(varName, typeName string) *VarValue {
	return &VarValue{
		name:     varName,
		typeName: newNameHelper("", typeName),
	}
}

// QualVar creates a new variable with a type name and its alias
func QualVar(varName, typeAlias, typeName string) *VarValue {
	return &VarValue{
		name:     varName,
		typeName: newNameHelper(typeAlias, typeName),
	}
}

// Pointer turns the variable type into a pointer
func (v *VarValue) Pointer() *VarValue {
	v.SetIsPointer(true)
	return v
}

// SetIsPointer sets whether or not a variable is a pointer
func (v *VarValue) SetIsPointer(isPointer bool) *VarValue {
	v.typeName.setIsPointer(isPointer)
	return v
}

func newVars(vars []*VarValue) *varsDeclarationStmt {
	return &varsDeclarationStmt{
		vars: vars,
	}
}

func (v *varsDeclarationStmt) write(sb *strings.Builder) {
	v.writeStmt(sb)
}

func (v *varsDeclarationStmt) writeStmt(sb *strings.Builder) bool {
	if len(v.vars) != 0 {
		typeMap := make(map[string][]string)
		keys := make([]string, 0)

		for _, val := range v.vars {
			typeName := val.typeName.getTypeName()

			if grouping, ok := typeMap[typeName]; ok {
				typeMap[typeName] = append(grouping, val.name)
			} else {
				keys = append(keys, typeName)
				typeMap[typeName] = []string{val.name}
			}
		}

		for _, typeName := range keys {
			vars := typeMap[typeName]

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

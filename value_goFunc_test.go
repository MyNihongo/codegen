package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoFuncLenField(t *testing.T) {
	const want = `len(a.field)`

	var sb strings.Builder
	Len(Identifier("a").Field("field")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncLenFuncCall(t *testing.T) {
	const want = `len(alias.myFunc())`

	var sb strings.Builder
	Len(QualFuncCall("alias", "myFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncEquals(t *testing.T) {
	const want = `len(a)==0`

	var sb strings.Builder
	Len(Identifier("a")).Equals(Int(0)).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncNotEquals(t *testing.T) {
	const want = `len(a)!=0`

	var sb strings.Builder
	Len(Identifier("a")).NotEquals(Int(0)).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

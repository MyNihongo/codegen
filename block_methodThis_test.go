package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncThis(t *testing.T) {
	const want = `t Type`

	var sb strings.Builder
	This("Type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncThisPointer(t *testing.T) {
	const want = `t *Type`

	var sb strings.Builder
	This("Type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualThis(t *testing.T) {
	const want = `t alias.Type`

	var sb strings.Builder
	QualThis("alias", "Type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualThisPointer(t *testing.T) {
	const want = `t *alias.Type`

	var sb strings.Builder
	QualThis("alias", "Type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestPanicEmptyTypeName(t *testing.T) {
	assert.Panics(t, func() {
		var sb strings.Builder
		This("").wr(&sb)
	})
}

package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncThis(t *testing.T) {
	const want = `t Type`

	var sb strings.Builder
	NewThis("Type").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncThisPointer(t *testing.T) {
	const want = `t *Type`

	var sb strings.Builder
	NewThis("Type").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualThis(t *testing.T) {
	const want = `t alias.Type`

	var sb strings.Builder
	NewQualThis("alias", "Type").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualThisPointer(t *testing.T) {
	const want = `t *alias.Type`

	var sb strings.Builder
	NewQualThis("alias", "Type").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

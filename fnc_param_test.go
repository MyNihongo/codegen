package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncParam(t *testing.T) {
	const want = `name type`

	var sb strings.Builder
	NewParam("name", "type").getValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamPointer(t *testing.T) {
	const want = `name *type`

	var sb strings.Builder
	NewParam("name", "type").Pointer().
		getValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualParam(t *testing.T) {
	const want = `name alias.type`

	var sb strings.Builder
	NewQualParam("name", NewQualName("alias", "type")).getValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualParamPointer(t *testing.T) {
	const want = `name *alias.type`

	var sb strings.Builder
	NewQualParam("name", NewQualName("alias", "type")).Pointer().
		getValue(&sb)

	assert.Equal(t, want, sb.String())
}

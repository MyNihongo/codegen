package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncParam(t *testing.T) {
	const want = `name type`

	var sb strings.Builder
	NewParam("name", "type").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamPointer(t *testing.T) {
	const want = `name *type`

	var sb strings.Builder
	NewParam("name", "type").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualParam(t *testing.T) {
	const want = `name alias.type`

	var sb strings.Builder
	NewQualParam("name", "alias", "type").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualParamPointer(t *testing.T) {
	const want = `name *alias.type`

	var sb strings.Builder
	NewQualParam("name", "alias", "type").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamsEmpty(t *testing.T) {
	const want = `()`

	var sb strings.Builder
	params := make([]*paramVal, 0)
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamsOne(t *testing.T) {
	const want = `(name1 type)`

	var sb strings.Builder
	params := []*paramVal{
		NewParam("name1", "type"),
	}
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncParams(t *testing.T) {
	const want = `(name1 type,name2 alias.type,name3 *type,name4 *alias.type)`

	var sb strings.Builder
	params := []*paramVal{
		NewParam("name1", "type"),
		NewQualParam("name2", "alias", "type"),
		NewParam("name3", "type").Pointer(),
		NewQualParam("name4", "alias", "type").Pointer(),
	}
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	const want = `obj`

	var sb strings.Builder
	Identifier("obj").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierPointer(t *testing.T) {
	const want = `*obj`

	var sb strings.Builder
	Identifier("obj").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifier(t *testing.T) {
	const want = `alias.obj`

	var sb strings.Builder
	QualIdentifier("alias", "obj").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierPointer(t *testing.T) {
	const want = `*alias.obj`

	var sb strings.Builder
	QualIdentifier("alias", "obj").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	const want = `new(myType)`

	var sb strings.Builder
	New("myType").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualNew(t *testing.T) {
	const want = `new(alias.MyType)`

	var sb strings.Builder
	QualNew("alias", "MyType").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNotPointer(t *testing.T) {
	got := New("myType").isPointer()
	assert.False(t, got)

	got = QualNew("alias", "MyType").isPointer()
	assert.False(t, got)
}

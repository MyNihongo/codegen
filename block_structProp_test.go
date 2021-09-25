package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProp(t *testing.T) {
	const want = `prop myType`

	var sb strings.Builder
	Property("prop", "myType").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEmbeddedProp(t *testing.T) {
	const want = `myType`

	var sb strings.Builder
	EmbeddedProperty("myType").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestPropPointer(t *testing.T) {
	const want = `prop *myType`

	var sb strings.Builder
	Property("prop", "myType").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEmbeddedPropPointer(t *testing.T) {
	const want = `*myType`

	var sb strings.Builder
	EmbeddedProperty("myType").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualProp(t *testing.T) {
	const want = `prop alias.myType`

	var sb strings.Builder
	QualProperty("prop", "alias", "myType").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualEmbeddedProp(t *testing.T) {
	const want = `alias.MyType`

	var sb strings.Builder
	QualEmbeddedProperty("alias", "MyType").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualPropPointer(t *testing.T) {
	const want = `prop *alias.myType`

	var sb strings.Builder
	QualProperty("prop", "alias", "myType").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualEmbeddedPropPointer(t *testing.T) {
	const want = `*alias.MyType`

	var sb strings.Builder
	QualEmbeddedProperty("alias", "MyType").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestPropSetIsPointerTrue(t *testing.T) {
	const want = `myProp *myType`

	var sb strings.Builder
	Property("myProp", "myType").SetIsPointer(true).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestPropSetIsPointerFalse(t *testing.T) {
	const want = `myProp myType`

	var sb strings.Builder
	Property("myProp", "myType").SetIsPointer(false).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

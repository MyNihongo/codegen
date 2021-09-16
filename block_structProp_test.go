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

func TestPropPointer(t *testing.T) {
	const want = `prop *myType`

	var sb strings.Builder
	Property("prop", "myType").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualProp(t *testing.T) {
	const want = `prop alias.myType`

	var sb strings.Builder
	QualProperty("prop", "alias", "myType").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualPropPointer(t *testing.T) {
	const want = `prop *alias.myType`

	var sb strings.Builder
	QualProperty("prop", "alias", "myType").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

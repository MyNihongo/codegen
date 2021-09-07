package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQualNameFull(t *testing.T) {
	const want = `alias.type`

	var sb strings.Builder
	NewQualName("alias", "type").getValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualName(t *testing.T) {
	const want = `type`

	var sb strings.Builder
	NewQualName("", "type").getValue(&sb)

	assert.Equal(t, want, sb.String())
}

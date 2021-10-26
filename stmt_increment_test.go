package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementIdentifier(t *testing.T) {
	const want = `i++`

	var sb strings.Builder
	got := newIncrement(Identifier("i")).writeStmt(&sb)

	assert.True(t, got)
	assert.Equal(t, want, sb.String())
}

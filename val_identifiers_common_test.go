package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	const want = `err`

	var sb strings.Builder
	Err().writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

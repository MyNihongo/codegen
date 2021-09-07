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

func TestNil(t *testing.T) {
	const want = `nil`

	var sb strings.Builder
	Nil().writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

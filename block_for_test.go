package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForNoCheckPanic(t *testing.T) {
	const want = `aaaa`

	var sb strings.Builder
	For(Declare("i").Values(Int(0)), nil, Identifier("i").Increment()).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

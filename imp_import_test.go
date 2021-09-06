package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePath(t *testing.T) {
	const want = `"path"
`

	var sb strings.Builder
	(&importStmt{path: "path"}).write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestWritePathAlias(t *testing.T) {
	const want = `alias "path"
`

	var sb strings.Builder
	(&importStmt{path: "path", alias: "alias"}).write(&sb)

	assert.Equal(t, want, sb.String())
}

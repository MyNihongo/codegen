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
	Import("path").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestWritePathAlias(t *testing.T) {
	const want = `alias "path"
`
	var sb strings.Builder
	ImportAlias("path", "alias").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEmptyGroup(t *testing.T) {
	var sb strings.Builder
	newImportsBlock().write(&sb)

	assert.Empty(t, sb.String())
}

func TestImportGroup(t *testing.T) {
	const want = `import "path"
`
	var sb strings.Builder
	fixture := newImportsBlock()
	fixture.AddImport("path")
	fixture.write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestImportAliasGroup(t *testing.T) {
	const want = `import alias "path"
`
	var sb strings.Builder
	fixture := newImportsBlock()
	fixture.AddImportAlias("path", "alias")
	fixture.write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestImportGroupMultiple(t *testing.T) {
	const want = `import (
alias "path"
"path2"
)
`
	var sb strings.Builder
	fixture := newImportsBlock()
	fixture.AddImportAlias("path", "alias")
	fixture.AddImport("path2")
	fixture.write(&sb)

	assert.Equal(t, want, sb.String())
}

package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGroup(t *testing.T) {
	var sb strings.Builder
	importGroup().write(&sb)

	assert.Empty(t, sb.String())
}

func TestImportGroup(t *testing.T) {
	const want = `import (
"path"
)
`
	var sb strings.Builder
	fixture := importGroup()
	fixture.append(&importBlock{path: "path"})
	fixture.write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestImportAliasGroup(t *testing.T) {
	const want = `import (
alias "path"
)
`
	var sb strings.Builder
	fixture := importGroup()
	fixture.append(&importBlock{path: "path", alias: "alias"})
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
	fixture := importGroup()
	fixture.append(&importBlock{path: "path", alias: "alias"})
	fixture.append(&importBlock{path: "path2"})
	fixture.write(&sb)

	assert.Equal(t, want, sb.String())
}

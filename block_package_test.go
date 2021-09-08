package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	const want = `package packageName
`
	var sb strings.Builder
	(&packageBlock{pkgName: "packageName"}).write(&sb)

	assert.Equal(t, want, sb.String())
}

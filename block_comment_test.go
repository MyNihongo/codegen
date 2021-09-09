package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentF(t *testing.T) {
	const want = `// this is a comment - 123
`

	var sb strings.Builder
	newCommentF("this is a %s - %d", "comment", 123).write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestComment(t *testing.T) {
	const want = `// this is a comment
`

	var sb strings.Builder
	newComment("this is a comment").write(&sb)

	assert.Equal(t, want, sb.String())
}

package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForNoCheckPanic(t *testing.T) {
	assert.Panics(t, func() {
		var sb strings.Builder
		For(Declare("i").Values(Int(0)), nil, Identifier("i").Increment()).
			write(&sb)
	})
}

func TestForOnlyCheck(t *testing.T) {
	const want = `for ;obj1.uid!=obj2.uid; {
}
`
	var sb strings.Builder
	For(nil, Identifier("obj1").Field("uid").NotEquals(Identifier("obj2").Field("uid")), nil).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestForInitAndCheck(t *testing.T) {
	const want = `for ;obj1.uid!=obj2.uid; {
}
`
	var sb strings.Builder
	For(Declare("i").Values(Int(0)), Identifier("i").Equals(Len(Identifier("myStr"))), nil).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

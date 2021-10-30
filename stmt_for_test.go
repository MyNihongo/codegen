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
	const want = `for i:=0;i<len(myStr); {
}
`
	var sb strings.Builder
	For(Declare("i").Values(Int(0)), Identifier("i").LessThan(Len(Identifier("myStr"))), nil).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestForCheckAndPost(t *testing.T) {
	const want = `for ;i<len(myStr);i++ {
}
`
	var sb strings.Builder
	For(nil, Identifier("i").LessThan(Len(Identifier("myStr"))), Identifier("i").Increment()).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestForAllStatements(t *testing.T) {
	const want = `for i:=0;i<len(myStr);i++ {
}
`
	var sb strings.Builder
	For(Declare("i").Values(Int(0)), Identifier("i").LessThan(Len(Identifier("myStr"))), Identifier("i").Increment()).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestForWithBlock(t *testing.T) {
	const want = `for true {
	if false {
		return
	}
}
`
	var sb strings.Builder
	For(nil, Identifier("true"), nil).Block(
		If(Identifier("false")).Block(
			Return(),
		),
	).write(&sb)

	assert.Equal(t, want, formatSb(sb))
}

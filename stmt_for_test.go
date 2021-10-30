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
			writeStmt(&sb)
	})
}

func TestForOnlyCheck(t *testing.T) {
	const want = `for ;obj1.uid!=obj2.uid; {
}
`
	var sb strings.Builder
	got := For(nil, Identifier("obj1").Field("uid").NotEquals(Identifier("obj2").Field("uid")), nil).
		writeStmt(&sb)

	assert.True(t, got)
	assert.Equal(t, want, sb.String())
}

func TestForInitAndCheck(t *testing.T) {
	const want = `for i:=0;i<len(myStr); {
}
`
	var sb strings.Builder
	got := For(Declare("i").Values(Int(0)), Identifier("i").LessThan(Len(Identifier("myStr"))), nil).
		writeStmt(&sb)

	assert.True(t, got)
	assert.Equal(t, want, sb.String())
}

func TestForCheckAndPost(t *testing.T) {
	const want = `for ;i<len(myStr);i++ {
}
`
	var sb strings.Builder
	got := For(nil, Identifier("i").LessThan(Len(Identifier("myStr"))), Identifier("i").Increment()).
		writeStmt(&sb)

	assert.True(t, got)
	assert.Equal(t, want, sb.String())
}

func TestForAllStatements(t *testing.T) {
	const want = `for i:=0;i<len(myStr);i++ {
}
`
	var sb strings.Builder
	got := For(Declare("i").Values(Int(0)), Identifier("i").LessThan(Len(Identifier("myStr"))), Identifier("i").Increment()).
		writeStmt(&sb)

	assert.True(t, got)
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
	got := For(nil, Identifier("true"), nil).Block(
		If(Identifier("false")).Block(
			Return(),
		),
	).writeStmt(&sb)

	assert.True(t, got)
	assert.Equal(t, want, formatSb(sb))
}

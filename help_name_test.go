package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQualNameFull(t *testing.T) {
	const want = `alias.type`

	var sb strings.Builder
	newNameHelper("alias", "type").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualName(t *testing.T) {
	const want = `type`

	var sb strings.Builder
	newNameHelper("", "type").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNameIsValidTrue(t *testing.T) {
	aliases, pointers := []string{"", "alias"}, []bool{true, false}

	for _, alias := range aliases {
		for _, pointer := range pointers {
			fixture := newNameHelper(alias, "type")
			fixture.setIsPointer(pointer)

			got := fixture.isValid()
			assert.True(t, got)
		}
	}
}

func TestNameIsValidFalse(t *testing.T) {
	aliases, pointers := []string{"", "alias"}, []bool{true, false}

	for _, alias := range aliases {
		for _, pointer := range pointers {
			fixture := newNameHelper(alias, "")
			fixture.setIsPointer(pointer)

			got := fixture.isValid()
			assert.False(t, got)
		}
	}
}

func TestNameIsArray(t *testing.T) {
	const want = `[]a`

	fixture := newNameHelper("", "a")
	fixture.setIsArray(true)

	var sb strings.Builder
	fixture.writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualNameIsArray(t *testing.T) {
	const want = `[]alias.MyType`

	fixture := newNameHelper("alias", "MyType")
	fixture.setIsArray(true)

	var sb strings.Builder
	fixture.writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNamePointerIsArray(t *testing.T) {
	const want = `[]*a`

	fixture := newNameHelper("", "a")
	fixture.setIsPointer(true)
	fixture.setIsArray(true)

	var sb strings.Builder
	fixture.writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualNamePointerIsArray(t *testing.T) {
	const want = `[]*alias.MyType`

	fixture := newNameHelper("alias", "MyType")
	fixture.setIsPointer(true)
	fixture.setIsArray(true)

	var sb strings.Builder
	fixture.writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

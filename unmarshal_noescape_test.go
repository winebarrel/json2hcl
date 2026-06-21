package json2hcl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/json2hcl"
)

func TestUnmarshalNoEscape(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// ${...} and %{...} are emitted literally instead of $${...} / %%{...}
	s, err := json2hcl.UnmarshalString(`{"foo":"_${bar}_","pct":"%{if x}"}`, json2hcl.NoEscape())
	require.NoError(err)
	assert.Equal("{\n  foo = \"_${bar}_\"\n  pct = \"%{if x}\"\n}", s)

	// a stray '${' is not mangled by formatting (no space inserted)
	s, err = json2hcl.UnmarshalString(`{"a":"${unclosed"}`, json2hcl.NoEscape())
	require.NoError(err)
	assert.Equal("{\n  a = \"${unclosed\"\n}", s)
}

func TestUnmarshalEscapeByDefault(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// without the option, template sequences are escaped
	s, err := json2hcl.UnmarshalString(`{"foo":"_${bar}_"}`)
	require.NoError(err)
	assert.Equal("{\n  foo = \"_$${bar}_\"\n}", s)
}

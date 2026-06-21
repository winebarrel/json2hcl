package json2hcl_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/json2hcl"
)

func TestUnmarshalError(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]string{
		"empty input":                "",
		"unterminated object":        "{",
		"unterminated array":         "[",
		"non-string object key":      "{1:2}",
		"missing object value":       `{"a":}`,
		"invalid array element":      "[}",
		"bad array element literal":  "[tru",
		"object missing close brace": `{"a":1`,
		"array missing close brack":  "[1",
	}

	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := json2hcl.Unmarshal([]byte(in))
			assert.Error(err)
		})
	}
}

func TestUnmarshalStringError(t *testing.T) {
	_, err := json2hcl.UnmarshalString("{")
	assert.Error(t, err)
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("read error")
}

func TestUnmarshalFromError(t *testing.T) {
	_, err := json2hcl.UnmarshalFrom(errReader{})
	assert.Error(t, err)
}

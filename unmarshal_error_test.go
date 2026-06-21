package json2hcl_test

import (
	"bytes"
	"encoding/json"
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

// The following cover defensive branches that the public API can't reach,
// because encoding/json rejects the corresponding inputs before they get this
// far. They exercise the existing code as-is rather than reshaping it.

func TestScalarTokensError(t *testing.T) {
	assert := assert.New(t)

	// json.Number always holds a valid number when produced by the decoder,
	// so this parse failure is only reachable by calling scalarTokens directly.
	_, err := json2hcl.ScalarTokens(json.Number("not-a-number"))
	assert.Error(err)

	// The decoder never yields a token of this type.
	_, err = json2hcl.ScalarTokens(123)
	assert.Error(err)
}

func TestDecodeValueUnexpectedDelim(t *testing.T) {
	assert := assert.New(t)

	dec := json.NewDecoder(bytes.NewReader([]byte("{}")))

	_, err := dec.Token() // consume the opening '{'
	assert.NoError(err)

	// The next token is the closing '}', a delimiter decodeValue never expects
	// in a value position.
	_, err = json2hcl.DecodeValue(dec)
	assert.Error(err)
}

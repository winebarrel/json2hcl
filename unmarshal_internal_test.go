package json2hcl

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// These cover defensive branches that the public API can't reach, because
// encoding/json rejects the corresponding inputs before they get this far.
// They exercise the existing code as-is rather than reshaping it for tests.

func TestScalarTokensError(t *testing.T) {
	assert := assert.New(t)

	// json.Number always holds a valid number when produced by the decoder,
	// so this parse failure is only reachable by calling scalarTokens directly.
	_, err := scalarTokens(json.Number("not-a-number"))
	assert.Error(err)

	// The decoder never yields a token of this type.
	_, err = scalarTokens(123)
	assert.Error(err)
}

func TestDecodeValueUnexpectedDelim(t *testing.T) {
	assert := assert.New(t)

	dec := json.NewDecoder(bytes.NewReader([]byte("{}")))

	_, err := dec.Token() // consume the opening '{'
	assert.NoError(err)

	// The next token is the closing '}', a delimiter decodeValue never expects
	// in a value position.
	_, err = decodeValue(dec)
	assert.Error(err)
}

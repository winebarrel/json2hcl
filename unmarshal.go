package json2hcl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// Option configures Unmarshal and its variants.
type Option func(*options)

type options struct {
	escapeTemplates bool
}

func defaultOptions() options {
	return options{escapeTemplates: true}
}

// NoEscape disables escaping of ${...} and %{...} template sequences, emitting
// the literal characters instead of $${...} / %%{...}. Note this makes those
// sequences behave as HCL template interpolations/directives rather than
// literal text.
func NoEscape() Option {
	return func(o *options) {
		o.escapeTemplates = false
	}
}

func Unmarshal(b []byte, opts ...Option) ([]byte, error) {
	o := defaultOptions()

	for _, fn := range opts {
		fn(&o)
	}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()

	toks, err := decodeValue(dec)

	if err != nil {
		return nil, err
	}

	// reject trailing data after the first JSON value
	if _, err := dec.Token(); err != io.EOF {
		if err == nil {
			err = fmt.Errorf("unexpected trailing data after JSON value")
		}

		return nil, err
	}

	out := bytes.TrimSpace(hclwrite.Format(toks.Bytes()))

	if !o.escapeTemplates {
		// Undo template-sequence escaping after formatting, so hclwrite.Format
		// still sees canonical (escaped) input and can't mangle stray '${'.
		out = bytes.ReplaceAll(out, []byte("$${"), []byte("${"))
		out = bytes.ReplaceAll(out, []byte("%%{"), []byte("%{"))
	}

	return out, nil
}

func UnmarshalString(s string, opts ...Option) (string, error) {
	bs, err := Unmarshal([]byte(s), opts...)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func UnmarshalFrom(r io.Reader, opts ...Option) ([]byte, error) {
	b, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return Unmarshal(b, opts...)
}

// decodeValue reads the next JSON value from dec and converts it into HCL
// tokens, preserving the order of object keys as they appear in the input.
func decodeValue(dec *json.Decoder) (hclwrite.Tokens, error) {
	tok, err := dec.Token()

	if err != nil {
		return nil, err
	}

	if delim, ok := tok.(json.Delim); ok {
		switch delim {
		case '{':
			return decodeObject(dec)
		case '[':
			return decodeArray(dec)
		default:
			return nil, fmt.Errorf("unexpected delimiter: %q", delim)
		}
	}

	return scalarTokens(tok)
}

func decodeObject(dec *json.Decoder) (hclwrite.Tokens, error) {
	var attrs []hclwrite.ObjectAttrTokens

	for dec.More() {
		keyTok, err := dec.Token()

		if err != nil {
			return nil, err
		}

		key := keyTok.(string)

		val, err := decodeValue(dec)

		if err != nil {
			return nil, err
		}

		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  nameTokens(key),
			Value: val,
		})
	}

	// consume the closing '}'
	if _, err := dec.Token(); err != nil {
		return nil, err
	}

	return hclwrite.TokensForObject(attrs), nil
}

func decodeArray(dec *json.Decoder) (hclwrite.Tokens, error) {
	elems := []hclwrite.Tokens{}

	for dec.More() {
		val, err := decodeValue(dec)

		if err != nil {
			return nil, err
		}

		elems = append(elems, val)
	}

	// consume the closing ']'
	if _, err := dec.Token(); err != nil {
		return nil, err
	}

	return tupleTokens(elems), nil
}

// tupleTokens builds an HCL tuple expression. A non-empty tuple is laid out
// with one element per line so that hclwrite.Format can indent it into a
// readable multi-line block, instead of the default single-line form.
func tupleTokens(elems []hclwrite.Tokens) hclwrite.Tokens {
	if len(elems) == 0 {
		return hclwrite.TokensForTuple(elems)
	}

	toks := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrack, Bytes: []byte{'['}},
		{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}},
	}

	for _, elem := range elems {
		toks = append(toks, elem...)
		toks = append(toks,
			&hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte{','}},
			&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}},
		)
	}

	toks = append(toks, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte{']'}})

	return toks
}

func scalarTokens(tok json.Token) (hclwrite.Tokens, error) {
	switch v := tok.(type) {
	case nil:
		return hclwrite.TokensForValue(cty.NullVal(cty.DynamicPseudoType)), nil
	case bool:
		return hclwrite.TokensForValue(cty.BoolVal(v)), nil
	case string:
		return hclwrite.TokensForValue(cty.StringVal(v)), nil
	case json.Number:
		nv, err := cty.ParseNumberVal(string(v))

		if err != nil {
			return nil, err
		}

		return hclwrite.TokensForValue(nv), nil
	default:
		return nil, fmt.Errorf("unsupported JSON value: %T", v)
	}
}

func nameTokens(key string) hclwrite.Tokens {
	if hclsyntax.ValidIdentifier(key) {
		return hclwrite.TokensForIdentifier(key)
	}

	return hclwrite.TokensForValue(cty.StringVal(key))
}

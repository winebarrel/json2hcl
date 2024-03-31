package json2hcl

import (
	"io"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty/json"
)

func Unmarshal(b []byte) ([]byte, error) {
	var v json.SimpleJSONValue
	err := (&v).UnmarshalJSON(b)

	if err != nil {
		return nil, err
	}

	toks := hclwrite.TokensForValue(v.Value)
	return toks.Bytes(), nil
}

func UnmarshalString(s string) (string, error) {
	bs, err := Unmarshal([]byte(s))

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func UnmarshalFrom(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return Unmarshal(b)
}

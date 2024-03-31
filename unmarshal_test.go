package json2hcl_test

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/json2hcl"
)

func subExt(t *testing.T, file string, newExt string) string {
	t.Helper()
	ext := path.Ext(file)
	return file[0:len(file)-len(ext)] + newExt
}

func TestUnmarshal(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	jsons, err := filepath.Glob("./tests/*.json")
	require.NoError(err)

	for _, j := range jsons {
		jsonStr, err := os.ReadFile(j)
		require.NoError(err)
		h := subExt(t, j, ".hcl")
		hclStr, err := os.ReadFile(h)
		require.NoError(err)

		fmt.Printf("%s -> %s\n", j, h)
		b, err := json2hcl.Unmarshal(jsonStr)
		require.NoError(err)
		assert.Equal(string(bytes.TrimSpace(hclStr)), string(b))
	}
}

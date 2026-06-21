package json2hcl

// Aliases that expose unexported helpers to the json2hcl_test package, so the
// external tests can exercise defensive branches the public API can't reach.
var (
	ScalarTokens = scalarTokens
	DecodeValue  = decodeValue
)

package main

import (
	"skywatch.danielnemeth.net/internal/assert"
	"testing"
)

func TestUrlParts_Compose(t *testing.T) {
	testUrl := urlParts{
		base:      "https://base-url.com",
		pathParam: "v2/some-api/provides/some-function",
	}
	actual := testUrl.Compose()
	expected := "https://base-url.com/v2/some-api/provides/some-function"

	assert.Equal(t, actual, expected)
}

func TestUrlParts_Compose2(t *testing.T) {
	testUrl := urlParts{
		base:      "https://base-url.com",
		pathParam: "v2/some-api/provides/some-function",
		airport:   "MAD",
	}
	actual := testUrl.Compose()
	expected := "https://base-url.com/v2/some-api/provides/some-function?airport=MAD"

	assert.Equal(t, actual, expected)
}

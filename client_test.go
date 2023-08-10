package main

import (
	"bytes"
	"skywatch.danielnemeth.net/internal/assert"
	"testing"
)

func TestSkyScannerClient_PollUntilCompletes(t *testing.T) {
	testBody := []byte("{\"status\":\"RESULT_STATUS_COMPLETE\",\"token\":\"test-token\"}")
	client := SkyScannerClient{
		apiKey: "test-apikey",
		urlParts: urlParts{
			base: "https://test/url",
		},
		PayloadBuilder: PayloadBuilder{},
	}
	result := client.PollUntilCompletes(testBody)
	assert.Equal(t, bytes.Equal(testBody, result), true)
}

func TestSkyScannerClient_PollUntilCompletesError(t *testing.T) {
	testBody := []byte("{\"status\":\"RESULT_STATUS_FAILED\",\"token\":\"test-token\"}")
	client := SkyScannerClient{
		apiKey: "test-apikey",
		urlParts: urlParts{
			base: "https://test/url",
		},
		PayloadBuilder: PayloadBuilder{},
	}
	result := client.PollUntilCompletes(testBody)
	assert.Equal(t, bytes.Equal(testBody, result), true)
}

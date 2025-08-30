package utils

import (
	"encoding/json"
	"io"
	"testing"
)

func CreateMapFromBody(responseBody io.ReadCloser, t *testing.T) map[string]any {
	jsonBytes, err := io.ReadAll(responseBody)
	if err != nil {
		t.Fatalf("Failed reading response body: %v", err)
	}

	var jsonBody map[string]any
	if err := json.Unmarshal(jsonBytes, &jsonBody); err != nil {
		t.Fatalf("Failed parsing response body: %v", err)
	}

	return jsonBody
}

package utils

import (
	"encoding/json"
	"io"
	"net/http"
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

func HTTPPut(url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}

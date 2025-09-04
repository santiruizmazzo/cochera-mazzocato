package utils

import (
	"maps"
	"testing"
)

func AssertResponseContains[C comparable](response map[string]any, expectedKey string, expectedContent C, t *testing.T) {
	if receivedContent, ok := response[expectedKey]; !ok || receivedContent != expectedContent {
		t.Fatalf("Expected %s %v, got %v", expectedKey, expectedContent, receivedContent)
	}
}

func AssertStatusCodeIs(expectedCode, receivedCode int, t *testing.T) {
	if expectedCode != receivedCode {
		t.Fatalf("Expected status code %d, got %d", expectedCode, receivedCode)
	}
}

func AssertSliceOfMaps(t *testing.T, data any) []map[string]any {
	// Manejar []map[string]any
	if slice, ok := data.([]map[string]any); ok {
		result := make([]map[string]any, len(slice))
		for i, m := range slice {
			converted := make(map[string]any)
			maps.Copy(converted, m)
			result[i] = converted
		}
		return result
	}

	// Manejar []any
	if slice, ok := data.([]any); ok {
		result := make([]map[string]any, len(slice))
		for i, item := range slice {
			if m, ok := item.(map[string]any); ok {
				result[i] = m
			} else if m, ok := item.(map[string]any); ok {
				converted := make(map[string]any)
				maps.Copy(converted, m)
				result[i] = converted
			} else {
				t.Fatalf("Element %d is not a map, got %T", i, item)
			}
		}
		return result
	}

	t.Fatalf("Data is not a slice of maps, got %T", data)
	return nil
}

package utils

import (
	"encoding/json"
	"testing"
)

func TestByteToMapping(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		wantLen   int
		wantKey   string
		wantValue string
	}{
		{
			name:      "valid JSON",
			input:     []byte(`{"key":"value","foo":"bar"}`),
			wantLen:   2,
			wantKey:   "key",
			wantValue: "value",
		},
		{
			name:    "empty JSON object",
			input:   []byte(`{}`),
			wantLen: 0,
		},
		{
			name:    "nil input",
			input:   nil,
			wantLen: 0,
		},
		{
			name:    "invalid JSON",
			input:   []byte(`{invalid}`),
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ByteToMapping[string](tt.input)
			if len(result) != tt.wantLen {
				t.Errorf("ByteToMapping() length = %d, want %d", len(result), tt.wantLen)
			}
			if tt.wantKey != "" {
				if result[tt.wantKey] != tt.wantValue {
					t.Errorf("ByteToMapping() key %s value = %s, want %s", tt.wantKey, result[tt.wantKey], tt.wantValue)
				}
			}
		})
	}
}

func TestMappingToByte(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]string
		wantValid bool
	}{
		{
			name:      "valid mapping",
			input:     map[string]string{"key": "value"},
			wantValid: true,
		},
		{
			name:      "empty mapping",
			input:     map[string]string{},
			wantValid: true,
		},
		{
			name:      "nil mapping",
			input:     nil,
			wantValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MappingToByte(tt.input)
			if result == nil {
				t.Errorf("MappingToByte() returned nil")
			}
			if len(result) == 0 {
				t.Errorf("MappingToByte() returned empty bytes")
			}
			// Verify it's valid JSON
			var m map[string]string
			if err := json.Unmarshal(result, &m); err != nil {
				t.Errorf("MappingToByte() result is not valid JSON: %v", err)
			}
		})
	}
}

func TestStructToMap(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name    string
		input   any
		wantLen int
	}{
		{
			name:    "valid struct",
			input:   User{Name: "John", Age: 30},
			wantLen: 2,
		},
		{
			name:    "empty struct",
			input:   User{},
			wantLen: 2,
		},
		{
			name:    "nil input",
			input:   nil,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StructToMap[any](tt.input)
			if len(result) != tt.wantLen {
				t.Errorf("StructToMap() length = %d, want %d", len(result), tt.wantLen)
			}
		})
	}
}

func TestStructToBytes(t *testing.T) {
	type Product struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price float64 `json:"price"`
	}

	tests := []struct {
		name      string
		input     any
		wantValid bool
	}{
		{
			name:      "valid struct",
			input:     Product{ID: 1, Name: "Widget", Price: 9.99},
			wantValid: true,
		},
		{
			name:      "empty struct",
			input:     Product{},
			wantValid: true,
		},
		{
			name:      "nil input",
			input:     nil,
			wantValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StructToBytes(tt.input)
			if result == nil {
				t.Errorf("StructToBytes() returned nil")
			}
			if len(result) == 0 {
				t.Errorf("StructToBytes() returned empty bytes")
			}
			// Verify it's valid JSON
			var m map[string]any
			if err := json.Unmarshal(result, &m); err != nil {
				t.Errorf("StructToBytes() result is not valid JSON: %v", err)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	// Test ByteToMapping → MappingToByte round trip
	original := map[string]string{"key1": "value1", "key2": "value2"}
	bytes := MappingToByte(original)
	result := ByteToMapping[string](bytes)

	if len(result) != len(original) {
		t.Errorf("Round trip length mismatch: got %d, want %d", len(result), len(original))
	}

	for k, v := range original {
		if result[k] != v {
			t.Errorf("Round trip value mismatch for key %s: got %s, want %s", k, result[k], v)
		}
	}
}

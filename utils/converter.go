package utils

import "encoding/json"

// ByteToMapping converts bytes to map[string]T using JSON unmarshaling.
// Returns an empty map if input is nil or unmarshaling fails.
func ByteToMapping[T any](input []byte) map[string]T {
	if input == nil {
		return map[string]T{}
	}
	var result map[string]T
	if err := json.Unmarshal(input, &result); err != nil {
		return map[string]T{}
	}
	return result
}

// MappingToByte converts map[string]T to bytes using JSON marshaling.
// Returns []byte(`{}`) if input is nil or marshaling fails.
func MappingToByte[T any](input map[string]T) []byte {
	if input == nil {
		return []byte(`{}`)
	}
	data, err := json.Marshal(input)
	if err != nil {
		return []byte(`{}`)
	}
	return data
}

// StructToMap converts a struct to map[string]T using JSON marshaling and unmarshaling.
// Returns an empty map if marshaling fails.
func StructToMap[T any](input any) map[string]T {
	data, err := json.Marshal(input)
	if err != nil {
		return map[string]T{}
	}
	return ByteToMapping[T](data)
}

// StructToBytes converts a struct to bytes using JSON marshaling.
// Returns []byte(`{}`) if marshaling fails.
func StructToBytes(input any) []byte {
	data, err := json.Marshal(input)
	if err != nil {
		return []byte(`{}`)
	}
	return data
}

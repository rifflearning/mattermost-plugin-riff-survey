package model

import (
	"encoding/json"
)

// TODO: Check if these methods can be generic

// StringArrayToByte returns a string array as a byte array
func StringArrayToByte(s []string) []byte {
	b, _ := json.Marshal(s)
	return b
}

// DecodeStringArrayFromByte tries to create a string array from a byte array
func DecodeStringArrayFromByte(b []byte) []string {
	var s []string
	if err := json.Unmarshal(b, &s); err != nil {
		return make([]string, 0)
	}
	return s
}

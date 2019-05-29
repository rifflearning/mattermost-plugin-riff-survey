package model

import (
	"encoding/json"
)

//TODO: Check if we can use GetBytes for storing other models in DB

// GetBytes returns an interface as a byte array
func GetBytes(s interface{}) []byte {
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

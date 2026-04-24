package json

import (
	"encoding/json"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MustMarshal(v any) []byte {
	data, _ := json.Marshal(v)
	return data
}

func Unmarshal(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}

func MustUnmarshal(data string, v any) {
	json.Unmarshal([]byte(data), v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
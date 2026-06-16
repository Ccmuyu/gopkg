package json

import (
	"encoding/json"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MustMarshal(v any) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

func Unmarshal(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}

func MustUnmarshal(data string, v any) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
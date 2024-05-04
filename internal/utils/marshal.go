package utils

import (
	"encoding/json"
)

func HardMarshal[T string | []byte](data any) T {
	marshaled, _ := json.Marshal(data)
	return T(marshaled)
}

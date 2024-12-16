package utils

import (
	"bytes"
	"encoding/gob"
)

func SerializeBoard(src [][]string) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(src)
	return buf.Bytes()
}

func DeserializeBoard(src []byte) (result [][]string, err error) {
	decoder := gob.NewDecoder(bytes.NewBuffer(src))
	err = decoder.Decode(&result)
	return result, err
}

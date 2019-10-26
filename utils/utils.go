package utils

import (
	"bytes"
	"encoding/json"
	"log"
)

func ToString(v interface{}) string {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return string(out)
}

func ToIoReader(v interface{}) *bytes.Reader {
	vBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return bytes.NewReader(vBytes)
}

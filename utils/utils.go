package utils

import (
	"encoding/json"
	"log"
)

func ToString(v interface{}) string {
	out, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return string(out)
}

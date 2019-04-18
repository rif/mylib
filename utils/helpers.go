package utils

import (
	"encoding/json"
	"log"
)

func ToJSON(o interface{}) string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Print(err)
	}
	return string(b)
}

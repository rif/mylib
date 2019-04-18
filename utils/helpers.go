package utils

import (
	"encoding/json"
	"log"
)

func ToJSON(o interface{}) string {
	var a []int
	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Print(err)
	}
	return string(b)
}

package mylib

import (
	"fmt"

	"github.com/rif/mylib/utils"
)

func GenerateMessage(name string, x int) string {
	return fmt.Sprintf("hello %s %d", name, x)
}

func ToNiceJSON(o interface{}) string {
	return "Nice json representation: " + utils.ToJSON(o)
}

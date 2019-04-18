package mylib

import (
	"fmt"

	"github.com/rif/mylib/utils"
)

func GenerateMessage(x int) string {
	return fmt.Sprintf("hello world %d", x)
}

func ToNiceJSON(o interface{}) string {
	return "Nice json representation: " + utils.ToJSON(o)
}

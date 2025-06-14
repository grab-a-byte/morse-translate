package main

import (
	"fmt"
	"morse/morse"
)

func main() {
	res, err := morse.Translate(".-")
	fmt.Printf("value '%s' or err of '%s' \n", res, err)

	str, err := morse.Translate("Help")
	fmt.Printf("value '%s' or err of '%s'\n", str, err)
}

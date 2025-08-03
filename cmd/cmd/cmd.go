package main

import (
	"flag"
	"fmt"
	"morse/morse"
	"os"
)

func main() {
	typePtr := flag.String("type", "", "Specifies the type of the requested translation")
	flag.Parse()
	input := flag.Arg(0)

	if *typePtr == "MorseCode" {
		res, err := morse.DecodeMorse(morse.MorseCode(input))
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(43)
		}

		fmt.Println(res)
		return
	} else if *typePtr == "Text" {
		res, err := morse.EncodeMorse(input)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(41)
		}

		fmt.Println(res)
		return
	}

	res, err := morse.Translate(input)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(100)
	}

	fmt.Println(res)
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("1. ASCII to char")
	fmt.Println("2. char to ASCII")
	var req string

	fmt.Println("What is your requirment: ")
	fmt.Scanln(&req)
	if req == "1" {
		c := 'A' // rune (characters in Go are represented using `rune` data type)
		asciiValue := int(c)
		fmt.Printf("Ascii Value of %c = %d\n", c, asciiValue)
	}
	if req == "2" {
		fmt.Println("Enter character value: ")
		var in int
		fmt.Scanln(&in)
		asciiValue := in
		character := rune(asciiValue)
		fmt.Printf("Character corresponding to Ascii Code %d = %c\n", asciiValue, character)
	}

}

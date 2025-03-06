package main

import (
	"fmt"
)

func main() {
	aString := "Hello World! €"
	fmt.Println("First byte", string(aString[0]))

	r := '€'
	fmt.Printf("As as string: %s and as a character :%c\n", r, r)
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}

	fmt.Println()

	for _, v := range aString {
		fmt.Printf("%c ", v)
	}
	fmt.Println()

}

package main

import (
	"fmt"
	"slices"
)

func main() {
	b := make([]byte, 12)
	fmt.Println("byte slice:", b)

	b = []byte("byte slice €")
	fmt.Println("byte slice:", b)

	fmt.Println(b)
	fmt.Printf("byte slice as text: %s\n", b)
	fmt.Println("byte slice as text:", string(b))
	fmt.Println("len b", len(b))
	slices.Delete(b,1,3)
	fmt.Println(b)
	fmt.Printf("byte slice as text: %s\n", b)
	fmt.Println("byte slice as text:", string(b))
	fmt.Println("len b", len(b))
}

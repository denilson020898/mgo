package main

import (
	"fmt"
)

func main() {
	a1 := []int{1}
	a2 := []int{-1, -2}
	a5 := []int{10,11,12,13,14}
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)

	copy(a1, a2)
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	copy(a1, a5)
	fmt.Println("a1", a1)
	fmt.Println("a5", a5)

	copy(a5, a2)
	fmt.Println("a5", a5)
	fmt.Println("a2", a2)
}

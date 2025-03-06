package main

import (
	"fmt"
)

func main() {
	slice := make([]byte, 3)

	arrayPtr := (*[3]byte)(slice)
	fmt.Println("print arrayPtr contents:", arrayPtr)
	fmt.Printf("Data type: %T\n", arrayPtr)
	fmt.Println("arrayPtr[0]:", arrayPtr[0])

	slice2 := []int{-1, -2, -3}
	array := [3]int(slice2)
	fmt.Println("print array contents:", array)
	fmt.Printf("Data type: %T\n", array)
	fmt.Println("array[0]:", array[0])
}

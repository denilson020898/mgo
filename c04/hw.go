package main

import (
	"fmt"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main () {
    PrintSlice([]int{1,2,23})
    PrintSlice([]string{"aa", "bb", "cc"})
    PrintSlice([]float32{12.2,2.1,-23})
    PrintSlice([]float64{1,2,23})
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an interget argument")
		return
	}

	index := arguments[1]

	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("using index", i)
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	if i > len(aSlice)-1 {
		fmt.Println("cannot delete element", i)
		return
	}

	// ... operator expand the slice
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("after 1st deletion", aSlice)

	if i > len(aSlice)-1 {
		fmt.Println("cannot delete element", i)
		return
	}

	aSlice[i] = aSlice[len(aSlice)-1]
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("after 2nd deletion", aSlice)
}

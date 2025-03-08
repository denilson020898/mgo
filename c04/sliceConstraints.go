package main

import (
	"fmt"
)

func f1[S interface{ ~[]E }, E interface{}](x S) int {
	return len(x)
}

func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}

func f3[S ~[]E, E any](x S) int {
	return len(x)
}

func main() {
	s := []AnotherInt{0, 1, 2}
	fmt.Println(AddElements(s))
}

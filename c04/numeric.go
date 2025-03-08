package main

import (
	"fmt"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func Same[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("4 = 3 is", Add(4.1, 3))
	fmt.Println("4.1 = 4.15 is", Add(4.1, 4.15))
}

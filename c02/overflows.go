package main

import (
	"fmt"
	"math"
)

func main() {
	i := math.MaxInt - 10

	for {
		if i == math.MaxInt {
			break
		}
		i = i + 1
	}
	fmt.Println("Max:", i)
	fmt.Println("Max overflow:", i+1)

	i = math.MinInt + 10
	for {
		if i == math.MinInt {
			break
		}
		i = i - 1
	}
	fmt.Println("Min:", i)
	fmt.Println("Max underflow:", i-1)
}

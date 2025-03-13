package main

import "fmt"

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	return
}

func main() {

}

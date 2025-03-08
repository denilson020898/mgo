package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)

	aMap = nil
	if aMap == nil {
		fmt.Println("nil map bro")
		aMap = map[string]int{}
	}
	_, ok := aMap["test"]
	if !ok {
		println("there is no aMap[\"test\"]")
	}

	aMap["test"] = 1
	v, ok := aMap["test"]
	if !ok {
		println("there is no aMap[\"test\"]")
	}
	println("there is aMap[\"test\"]:", v)
}

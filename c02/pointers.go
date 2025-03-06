package main

import (
	"fmt"
)

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123

	fmt.Println("memory addess of f:", &f)

	fP := &f
	fmt.Println("memory addess of fP:", &fP)
	fmt.Println("value of fP:", *fP)

	processPointer(fP)
	fmt.Printf("value of f: %.2f\n", f)

	x := returnPointer(f)
	fmt.Printf("value of x: %.2f\n", *x)

	xx := bothPointers(fP)
	fmt.Printf("value of xx: %.2f\n", *xx)

	var k *aStructure
	fmt.Println(k)
	if k == nil {
		k = new(aStructure)
	}
	fmt.Println(k)
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}

}

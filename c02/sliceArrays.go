package main

import (
	"fmt"
)

func change(s []string) {
	s[0] = "change_function"
}

func bar(slice []int) int {
	a := (*[3]int)(slice)
	return a[0] + a[1] + a[2] + a[3]
}

func main() {
	a := [4]string{"zero", "one", "two", "three"}
	fmt.Println("a:", a)

	S0 := a[0:1]
	fmt.Println(S0)
	S0[0] = "S0"

	S12 := a[1:3]

	fmt.Println(S12)
	S12[0] = "S12_0"
	S12[1] = "S12_1"

	fmt.Println("a:", a)
	change(S12)

	fmt.Println("a:", a)

	fmt.Println("capacity of S0:", cap(S0), "len of S0:", len(S0))
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"

	fmt.Println("a:", a)
}

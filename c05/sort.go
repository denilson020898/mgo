package main

import (
	"fmt"
	"sort"
)

type Size struct {
	F1 int
	F2 string
	F3 int
}

type Person struct {
	F1 int
	F2 string
	F3 Size
}

type Personslice []Person

func (ps Personslice) Len() int {
	return len(ps)
}
func (ps Personslice) Less(i, j int) bool {
	return ps[i].F3.F1 < ps[j].F3.F1
}

func (ps Personslice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	data := []Person{
		Person{1, "One", Size{1, "Person_1", 10}},
		Person{2, "One", Size{2, "Person_2", 20}},
		Person{-1, "One", Size{-1, "Person_3", -20}},
	}

	fmt.Println("Before:", data)
	sort.Sort(Personslice(data))
	fmt.Println("After:", data)

	sort.Sort(sort.Reverse(Personslice(data)))
	fmt.Println("Reversed:", data)
}

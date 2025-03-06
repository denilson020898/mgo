package main

import "fmt"

func main() {
	aSlice := []float64{}

	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	fmt.Println(aSlice, "with length", len(aSlice))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	t = append(t, -5)
	fmt.Println(t, len(t), cap(t))
	fmt.Println(t, "with length", len(t))

	twoD := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for _, i := range twoD {
		for _, j := range i {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4, 5}
	make2D[1] = []int{-1, -2, -3, -4, -5}
	fmt.Println(make2D)
}

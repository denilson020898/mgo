package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))

	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}
	return normalized
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	arguments := os.Args

	var values []float64
	if len(arguments) == 1 {
		values = make([]float64, 10)
		for i := 0; i < len(values); i++ {
			values[i] = randomFloat(-10, 10)
		}
	} else {
		values = make([]float64, len(arguments)-1)
		for i := 1; i < len(arguments); i++ {
			n, err := strconv.ParseFloat(arguments[i], 64)
			if err != nil {
				fmt.Printf("Invalid input: unable to parse %s\n", n)
				return
			}
			values[i-1] = n
		}
	}

	var min, max float64
	var initialized = 0
	var sum float64

	for _, n := range values {
		sum = sum + n
		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

	}

	nValues := len(values)

	fmt.Println("Number of values:", nValues)
	fmt.Println("Min", min)
	fmt.Println("Max", max)

	meanValue := sum / float64(nValues)
	fmt.Printf("Mean value: %.5f\n", meanValue)

	var squared float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		squared = squared + math.Pow((n-meanValue), 2)
	}
	standardDeviation := math.Sqrt(squared / float64(nValues))
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	normalizedValues := normalize(values, meanValue, standardDeviation)
	fmt.Println("Original:", values)
	fmt.Println("Normalized:", normalizedValues)
}

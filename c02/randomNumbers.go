package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func randomWithSeed(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	args := os.Args

	var minInt, maxInt, numInt, seed int
	var seeded bool

	if len(args) == 1 {
		minInt = 0
		maxInt = 99
		numInt = 100
		fmt.Println("Using default values!")
		seeded = false
	} else {
		fmt.Println("Arguments are detected!")

		n, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("can't parse %s\n", args[1])
			return
		}
		minInt = n

		n, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("can't parse %s\n", args[2])
			return
		}
		maxInt = n

		n, err = strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("can't parse %s\n", args[3])
			return
		}
		numInt = n

		n, err = strconv.Atoi(args[4])
		if err != nil {
			fmt.Printf("can't parse %s\n", args[4])
			return
		}
		seed = n
		seeded = true
	}

	rng := rand.New(rand.NewSource(int64(seed)))

	for i := 0; i < numInt; i++ {
        var val int
		if seeded {
			val = randomWithSeed(minInt, maxInt, rng)
		} else {
			val = random(minInt, maxInt)
		}
		fmt.Printf("%d ", val)
	}

}

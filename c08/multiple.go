package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need 1 interger argument")
		return
	}

	countVar := os.Args[1]
	count, err := strconv.Atoi(countVar)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}

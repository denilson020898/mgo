package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
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

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()

	fmt.Println("\nExiting...")
}

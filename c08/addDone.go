package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	flag := true

	if len(os.Args) == 1 {
		flag = false
	}

	count := 20

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

	fmt.Println("Add more waitGroup.Done()")
	if flag {
		waitGroup.Add(1)
	} else {
		waitGroup.Done()
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()

	fmt.Println("\nExiting...")
}

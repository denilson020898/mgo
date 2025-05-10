package main

import (
	"fmt"
)

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main() {
	var ch chan bool = make(chan bool)

	go printer(ch, 5)

	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()

}

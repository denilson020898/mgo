package main

import "fmt"

func main() {
	var c chan string

	c = make(chan string, 4)

	go func() {
		for i := 0; i < 50; i++ {
			c <- "hello"
		}
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
}

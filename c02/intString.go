package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Println("Need more argument")
		return
	}
	
	// n, err := strconv.ParseInt(arguments[1], 10, 10)
	n, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Println("invalid input %s\n", arguments[1])
		return
	}

	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)
	input = string(n)
	fmt.Printf("string() %s of type %T\n", input, input)
}

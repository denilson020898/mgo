package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	args := os.Args

	length := 8

	if len(args) != 1 {
		fmt.Println()
		var err error
		length, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("unable to parse %s\n", args[1])
			return
		}
	} else {
		fmt.Println("Using default values...")
	}

	password, err := generateBytes(int64(length))
	if err != nil {
		fmt.Println("generate crypto rand went wrong")
		return
	}

	fmt.Println(password)

}

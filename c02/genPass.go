package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	MIN int = 0
	MAX int = 94
)


func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
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

	password := getString(int64(length))

	fmt.Println(password)
	
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need more input in __%s__ format!", time.DateTime)
		return
	}

	input := os.Args[1]
	// now, err := time.Parse("02 January 2006 15:04 MST", input)
	now, err := time.Parse(time.DateTime, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// local time
	loc, _ := time.LoadLocation("Local")
	fmt.Printf("Current Location: %s\n", now.In(loc))
	
	// NY
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	// London
	loc, _ = time.LoadLocation("Europe/London")
	fmt.Printf("London Time: %s\n", now.In(loc))

	// Tokyo
	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo Time: %s\n", now.In(loc))

	// Jakarta
	loc, _ = time.LoadLocation("Asia/Jakarta")
	fmt.Printf("Jakarta Time: %s\n", now.In(loc))
}

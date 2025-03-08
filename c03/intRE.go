package main

import (
	"fmt"
	"regexp"
	"os"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need arguments")
		return
	}
	fmt.Println(matchInt(os.Args[1]))
}

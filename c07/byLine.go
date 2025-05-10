package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByline(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\r')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}

		if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need one argument")
		return
	}
	err := lineByline(os.Args[1])
	if err != nil {
		fmt.Println("can't read file line by line")
		return
	}
}

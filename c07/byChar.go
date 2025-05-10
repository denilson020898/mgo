package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {

			if len(line) != 0 {
				for _, x := range line {
					fmt.Println(string(x))
				}
			}
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need one argument")
		return
	}
	err := charByChar(os.Args[1])
	if err != nil {
		fmt.Println("can't read file line by line")
		return
	}
}

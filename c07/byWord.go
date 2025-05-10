package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByword(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	re := regexp.MustCompile("[^\\s]+")
	for {
		line, err := r.ReadString('\n')
		if err == nil {
			if len(line) != 0 {
				words := re.FindAllString(line, -1)
				for i := 0; i < len(words); i++ {
					fmt.Println(words[i])
				}
			}

		} else if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		// fmt.Print(line)
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need one argument")
		return
	}
	err := wordByword(os.Args[1])
	if err != nil {
		fmt.Println("can't read file line by line")
		return
	}
}

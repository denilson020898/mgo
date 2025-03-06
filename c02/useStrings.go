package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("To Upper: %s\n", s.ToUpper("Hello THERE"))
	f("To Lower: %s\n", s.ToLower("Hello THERE"))

	f("%s\n", s.Title("tHis wiLl bE a TiTle!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "Is"))

	t := s.Fields("This is a string!")
	f("Fields, %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields, %v\n", len(t))

	f("%s\n", s.Split("abcde efg", ""))
	f("%s\n", s.Replace("abcde efg", "", "_", -1))
	f("%s\n", s.Replace("abcde efg", "", "_", 4))
	f("%s\n", s.Replace("abcde efg", "", "_", 2))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t", trimFunction))
}

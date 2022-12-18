package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		d := Rot13(args[0])
		for _, b := range d {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')

	}
}

func Rot13(s string) string {
	str := []rune(s)
	for i := range str {
		if str[i] == 'a' || str[i] == 'A' {
			str[i] = str[i] + 25
		}
	}

	return string(str)
}

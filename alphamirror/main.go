package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		a := alpha(args[0])
		for _, v := range a {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func alpha(s string) string {
	str := []rune(s)
	for i := range str {
		if str[i] == 'a' {
			str[i] = str[i] + 25
		}
		if str[i] == 'b' {
			str[i] = str[i] + 23
		}
		if str[i] == 'c' {
			str[i] = str[i] + 21
		}
		if str[i] == 'd' {
			str[i] = str[i] + 19
		}
		if str[i] == 'e' {
			str[i] = str[i] + 17
		}
		if str[i] == 'f' {
			str[i] = str[i] + 15
		}
		if str[i] == 'g' {
			str[i] = str[i] + 13
		}
		if str[i] == 'h' {
			str[i] = str[i] + 11
		}
		if str[i] == 'i' {
			str[i] = str[i] + 9
		}
		if str[i] == 'j' {
			str[i] = str[i] + 7
		}
		if str[i] == 'k' {
			str[i] = str[i] + 5
		}
		if str[i] == 'l' {
			str[i] = str[i] + 3
		}
		if str[i] == 'm' {
			str[i] = str[i] + 1
		}
		if str[i] == 'n' {
			str[i] = str[i] - 1
		}
		if str[i] == 'o' {
			str[i] = str[i] - 3
		}
		if str[i] == 'p' {
			str[i] = str[i] - 5
		}
		if str[i] == 'q' {
			str[i] = str[i] - 7
		}
		if str[i] == 'r' {
			str[i] = str[i] - 9
		}
		if str[i] == 's' {
			str[i] = str[i] - 11
		}
		if str[i] == 't' {
			str[i] = str[i] - 13
		}
		if str[i] == 'u' {
			str[i] = str[i] - 15
		}
		if str[i] == 'v' {
			str[i] = str[i] - 17
		}
		if str[i] == 'w' {
			str[i] = str[i] - 19
		}
		if str[i] == 'x' {
			str[i] = str[i] - 21
		}
		if str[i] == 'y' {
			str[i] = str[i] - 23
		}
		if str[i] == 'z' {
			str[i] = str[i] - 25
		}

	}
	return string(str)
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	} else if Atoi(arg[0]) == 0 {
		return
	}

	for i := 1; i <= 9; i++ {
		result := i * Atoi(string(arg[0]))
		z01.PrintRune(rune(i) + 48)
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		printstr(Atoi(string(arg[0])))
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printstr(result)
		z01.PrintRune('\n')

	}
}

func Atoi(s string) int {
	mult := 1
	res := 0
	for _, v := range s {
		if v < 0 {
			return 0
		}
		res *= mult
		a := int(v - 48)
		res += a
		mult = 10
	}
	return res
}

func printstr(a int) {
	b := []rune{}
	for a > 0 {
		q := a / 10
		r := a % 10
		b = append(b, rune(r+48))
		a = q
	}
	for i := len(b) - 1; i >= 0; i-- {
		z01.PrintRune(b[i])
	}
}

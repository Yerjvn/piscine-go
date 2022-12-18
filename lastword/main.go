package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}

	word := ""
	save := []string{}

	for _, b := range arg[0] {
		if b != ' ' {
			word = word + string(b) // WORD = "" + "w"
		} else {
			if word != "" {
				save = append(save, word)
				word = ""
			}
		}
	}
	if word != "" {
		save = append(save, word)
	}

	if len(save) == 0 {
		return
	}
	last := save[len(save)-1]

	for _, v := range last {
		z01.PrintRune(v)
	}
}

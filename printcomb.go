package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for g := j + 1; g <= '9'; g++ {
				if i < j && j < g {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(g)
					if i == '7' && j == '8' && g == '9' {
						z01.PrintRune(10)
					} else {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}

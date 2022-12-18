package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			f := b + 1
			for k := a; k <= '9'; k++ {
				for ; f <= '9'; f++ {
					z01.PrintRune((a))
					z01.PrintRune((b))
					z01.PrintRune(' ')
					z01.PrintRune((k))
					z01.PrintRune((f))
					if a < '9' || b < '8' || k < '9' || f < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				f = '0'
			}
		}
	}
	z01.PrintRune('\n')
}

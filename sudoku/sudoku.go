package main

import (
	"os"

	"github.com/01-edu/z01"
)

var field [9][9]rune

func main() {
	if len(os.Args) != 10 {
		Error()
		return
	}
	InsertField()
	if checking() {
		if sudokusolver(0, 0) {
			output(field)
		} else {
			Error()
		}
	
	} else {
		Error()
	}
}

func checking() bool {
	t := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (field[i][j] <= '9' && field[i][j] > '0') || field[i][j] == '.' {
				t = true
			} else {
				return false
			}
		}
	}
	return t
}

func checkver(x, y, v int) bool {
	for i := 0; i < 9; i++ {
		if field[i][x] == rune(v)+48 {
			return true
		}
	}

	return false
}

func checkgor(x, y, v int) bool {
	for i := 0; i < 9; i++ {
		if field[y][i] == rune(v)+48 {
			return true
		}
	}

	return false
}

func InsertField() bool {
	rows := os.Args[1:]
	for i := 0; i < 9; i++ {
		if len(rows[i]) != 9 {
			return false
		}
		row := []rune(rows[i])
		for j := 0; j < 9; j++ {
			field[i][j] = row[j]
		}
	}
	return true
}

func output(result [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(result[i][j])
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func check3x3(x, y, v int) bool {
	sx, sy := (x/3)*3, (y/3)*3

	for valuey := range [3]int{} {
		for valuex := range [3]int{} {
			if field[sy+valuey][sx+valuex] == rune(v)+48 {
				return true
			}
		}
	}

	return false
}

func checkallparametry(x, y, v int) bool {
	return !checkver(x, y, v) && !checkgor(x, y, v) && !check3x3(x, y, v)
}

func next(x, y int) (int, int) {
	xzibit, yzibit := (x+1)%9, y
	if xzibit == 0 {
		yzibit = y + 1
	}
	return xzibit, yzibit
}

func sudokusolver(x, y int) bool {
	ervalue := 0
	if y == 9 {
		return true
	}
	if field[y][x] != '.' {
		ervalue++
		return sudokusolver(next(x, y))
	} else {
		var b [9]rune
		for i := range b {
			v := i + 1
			if checkallparametry(x, y, v) {
				field[y][x] = rune(v) + 48
				if sudokusolver(next(x, y)) {
					return true
				}
				field[y][x] = '.'
			}
		}
		return false
	}
}

func Error() {
	z01.PrintRune('E')
	z01.PrintRune('r')
	z01.PrintRune('r')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune('\n')
}

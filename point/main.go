package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)

	runeX := []rune{'x', ' ', '=', ' '}
	runeY := []rune{'y', ' ', '=', ' '}

	x := points.x
	y := points.y

	xS := []int{}
	yS := []int{}

	for x > 0 {
		q := x / 10
		r := x % 10
		x = q
		xS = append(xS, r)
	}

	for y > 0 {
		q := y / 10
		r := y % 10
		y = q
		yS = append(yS, r)

	}

	for i := len(xS) - 1; i >= 0; i-- {
		runeX = append(runeX, rune(48+xS[i]))
	}
	for i := len(yS) - 1; i >= 0; i-- {
		runeY = append(runeY, rune(48+yS[i]))
	}

	for _, v := range runeX {
		z01.PrintRune(v)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for _, v := range runeY {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

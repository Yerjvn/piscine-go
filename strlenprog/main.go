package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Helloasdasd World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	length := 0
	for index := range s {
		length = index + 1
	}
	// for range s {

	// 	length++
	// }
	return length
}

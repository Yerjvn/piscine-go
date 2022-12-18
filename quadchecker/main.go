package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// ./quadA 4 3 | go run .
// o--o\n|  |\no--o\n

func main() {
	first_output, _ := ioutil.ReadAll(os.Stdin) // []bytes (string)
	x, y := 0, 0
	for _, value := range first_output {
		if value == byte('\n') {
			y++
		} else {
			x++ // fullSUM = x * y
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	x = x / y // x = 12/3 y=3 x=4
	massQuad := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	flag := false
	for _, value := range massQuad {
		cmd := exec.Command("./"+value, strconv.Itoa(x), strconv.Itoa(y)) // ./quadA 4 3
		output, _ := cmd.Output()                                         // o--o\n|  |\no--o\n
		if string(first_output) == string(output) && flag {
			fmt.Print(" || ")
			fmt.Printf("[%v] [%d] [%d]", value, x, y)
		}
		if string(first_output) == string(output) && !flag {
			fmt.Printf("[%v] [%d] [%d]", value, x, y) // PrintFormat %v, %d  [quadA] [4] [3] || [quadC] [4] [3]
			flag = true
		}
	}
	if !flag {
		fmt.Print("Not a Raid function")
	}
	fmt.Println()
}

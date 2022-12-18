package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args
	for i := range argument {
		if argument[i] == "01" || argument[i] == "galaxy" || argument[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}

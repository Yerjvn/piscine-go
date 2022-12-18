package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	} else if len(arg) == 0 {
		fmt.Println("File name missing")
		return
	}
	content, err := ioutil.ReadFile(arg[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(content))
}

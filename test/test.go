package main

import ( 
	"fmt" 
	"os"

)

func main() {
    args := os.Args

	b := args[len(args)-1]

	fmt.Println(b)
} 
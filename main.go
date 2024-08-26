package main

import (
	fn "asciiart/functions"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments")
	}
	fn.MyPrint()
}

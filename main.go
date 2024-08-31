package main

import (
	art "asciiart/functions"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments")
		return
	}

	asciiArtMap := art.AssignArt("standard.txt")
	art.AsciiPrint(args[1], asciiArtMap)
}

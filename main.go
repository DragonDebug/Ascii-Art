package main

import (
	art "asciiart/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments")
		return
	}

	input := args[1]
	if input == "" {
		return
	}

	// Creating the map for the ASCII art
	asciiArtMap := art.AssignArt("standard.txt")

	// Checking if the input consists of only ASCII characters
	if !art.IsAsciiInput(input, asciiArtMap) {
		fmt.Println("only ASCII characters are accepted as input")
		return
	}

	// spliting the input wherever a new line (\n) is detected
	inputArr := strings.Split(input, "\\n")

	// Printing the inputArr
	art.AsciiPrintArr(inputArr, asciiArtMap)
}

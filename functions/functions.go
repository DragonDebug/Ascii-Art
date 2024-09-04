package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

// AssignArt reads a file containing ASCII art and assigns each character's art to a map.
func AssignArt(banner string) map[rune][]string {
	// Opening the file specified by the 'banner' argument.
	file, err := os.Open(banner)
	if err != nil {
		fmt.Printf("error opening the file: %s", banner)
		os.Exit(1)
	}
	defer file.Close()

	// Create a scanner to read the file line by line.
	fileLines := bufio.NewScanner(file)

	// Initialize a map to store ASCII art for each character.
	asciiArtMap := make(map[rune][]string)
	var char rune = ' '
	var charArt []string

	// Iterate through each line in the file.
	for fileLines.Scan() {
		if len(fileLines.Text()) != 0 {
			charArt = append(charArt, fileLines.Text())
			// Assign the ASCII art to the character in the map when it reaches 8 lines.
			if len(charArt) == 8 {
				asciiArtMap[char] = charArt
				char++
				charArt = nil
			}
		}
	}
	return asciiArtMap
}

// Checking if the user input consists of only ASCII characters by checking if the key exists inside of the map
func IsAsciiInput(word string, asciiArtMap map[rune][]string) bool {
	for _, char := range word {
		_, exists := asciiArtMap[char]
		if exists {
			continue
		} else {
			return false
		}
	}
	return true
}

// Printing the ASCII art for a word
func AsciiPrint(word string, asciiArtMap map[rune][]string) {
	if word == "" {
		return
	}
	for i := 0; i <= 7; i++ {
		for _, char := range word {
			fmt.Print(asciiArtMap[char][i])
		}
		fmt.Println()
	}
}

// Printing the ASCII art from an array of strings using AsciiPrint
func AsciiPrintArr(inputArr []string, asciiArtMap map[rune][]string) {
	for _, input := range inputArr {
		if input == "" {
			fmt.Println()
		} else {
			AsciiPrint(input, asciiArtMap)
		}
	}
}

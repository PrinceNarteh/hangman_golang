package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	word         = "Secret"
	allow_errors = 7
	guesses      = []rune{'e'}
	done         = false
	reader       = bufio.NewReader(os.Stdin)
)

func main() {

	for !done {
		for _, letter := range word {
			if contains(letter, guesses) {
				fmt.Print(string(letter))
			} else {
				fmt.Print("_")
			}
		}

		char := getUserInput()
		guesses = append(guesses, char)
		if !contains(char, []rune(word)) {
			allow_errors -= 1
			if allow_errors == 0 {
				break
			}
		}

		done = true
	}
}

func contains(str rune, arr []rune) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func getUserInput() rune {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.ToLower(input)
	firstCharacter := input[0]
	return rune(firstCharacter)
}

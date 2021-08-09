package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	word         = "secret"
	allow_errors = 7
	guesses      = []rune{}
	done         = false
	reader       = bufio.NewReader(os.Stdin)
)

func main() {

	for !done {
		fmt.Print("==> ")
		for _, letter := range word {
			if contains(letter, guesses) {
				fmt.Print(string(letter))
			} else {
				fmt.Print("_")
			}
		}

		fmt.Println("")

		char := getUserInput()
		guesses = append(guesses, char)
		if !contains(char, []rune(word)) {
			allow_errors -= 1
			if allow_errors == 0 {
				break
			}
		}

		done = true
		for _, v := range word {
			if !contains(v, guesses) {
				done = false
			}
		}
	}

	if done {
		fmt.Println("")
		fmt.Println("Bravo!!!")
		fmt.Printf("You found the word! It was %v", word)
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Printf("Game Over! The word was %v", word)
		fmt.Println("")
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
	fmt.Println(fmt.Sprintf("Allow Error left %v.", allow_errors))
	fmt.Print("Next Guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.ToLower(input)
	firstCharacter := input[0]
	fmt.Println("---------------------")
	return rune(firstCharacter)
}

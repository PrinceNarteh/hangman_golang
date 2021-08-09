package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	word = "Secret"
	allow_errors = 7
	guesses = []rune{'e'}
	done = false
)

func main() {
	fmt.Println(done)
	
	for !done {
		for _, letter := range word {
			if contains(letter, &guesses) {
				fmt.Print(string(letter))
			} else {
				fmt.Print("_")
			}
		}
		done = true
		
		char := getUserInput()
		guesses = append(guesses, char)
		fmt.Println(guesses)
	}
}

func contains(str rune, arr *[]rune) bool {
	for _, v := range *arr {
		if v == str {
			return true
		}
	}
	return false 
}

func getUserInput() rune {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err !=nil {
		fmt.Println(err)
	}
	input = strings.Replace(input, "\n", "", -1)
	firstCharacter := input[0]
	return rune(firstCharacter)
}
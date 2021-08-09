package main

import (
	"fmt"
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
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func filtering(word string) string {

	var filterd string

	for _, char := range word {

		if unicode.IsLetter(char) {
			filterd += string(unicode.ToUpper(char))
		}

	}
	return filterd
}

func checkPalindrome(word string) bool {
	var backward int = len(word) - 1

	for forward := 0; forward < backward; forward++ {
		if word[forward] != word[backward] {
			return false
		}
		backward--
	}
	return true
}

func main() {
	fmt.Print("Enter a word : ")
	input := bufio.NewReader(os.Stdin)
	word, _ := input.ReadString('\n')
	word = strings.TrimSpace(word)

	filter := filtering(word)
	fmt.Print(checkPalindrome(filter))

}

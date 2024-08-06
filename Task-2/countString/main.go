package main

import (
	"fmt"
	"strings"
	"unicode"
	"bufio"
	"os"

)

func WordCount(word string)map[string]int{

	count := make(map[string]int)

	for _, char := range word{

		if unicode.IsPunct(char) || unicode.IsSpace(char){
			continue
		}

		if unicode.IsLetter(char){
			character :=strings.ToUpper(string(char))
			count[character] += 1
		}else{
			count[string(char)] += 1
		}
		
	}
	return count
}

func Printing(stored map[string]int){

		for char, count := range stored {
			fmt.Printf("%s: %.d ", char, count)
	}
}
func main(){

	fmt.Print("Enter a word : ")
    input := bufio.NewReader(os.Stdin)
    word, _ := input.ReadString('\n')
    word = strings.TrimSpace(word) 

	counted := WordCount(word)
	Printing(counted)
}
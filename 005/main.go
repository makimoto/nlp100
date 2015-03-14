package main

import (
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	const input = "I am an NLPer"
	fmt.Println("Character bi-grams:")
	pp.Println(charBiGrams(input))
	fmt.Println("Word bi-grams:")
	pp.Println(wordBiGrams(input))
}

func charBiGrams(s string) []string {
	runes := []rune(s)
	var output []string
	for i := 0; i < len(runes)-1; i++ {
		output = append(output, string(runes[i:i+2]))
	}
	return output
}

func wordBiGrams(s string) [][]string {
	words := strings.Fields(s)
	var output [][]string
	for i := 0; i < len(words)-1; i++ {
		output = append(output, words[i:i+2])
	}
	return output
}

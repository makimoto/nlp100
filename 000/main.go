package main

import "fmt"

func main() {
	input := "stressed"
	runes := []rune(input)
	lenRunes := len(runes)

	for i := 0; i < lenRunes/2; i++ {
		runes[i], runes[lenRunes-i-1] = runes[lenRunes-i-1], runes[i]
	}

	fmt.Printf("%s", string(runes))
}

package main

import "fmt"

func main() {
	input := "パタトクカシーー"
	runes := []rune(input)
	output := append([]rune{}, runes[0], runes[2], runes[4], runes[6])
	fmt.Printf("%s", string(output))
}

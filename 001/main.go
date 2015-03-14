package main

import "fmt"

func main() {
	input := "パタトクカシーー"
	runes := []rune(input)
	output := append([]rune{}, runes[1], runes[3], runes[5], runes[7])
	fmt.Printf("%s", string(output))
}

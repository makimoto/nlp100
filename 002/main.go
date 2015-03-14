package main

import "fmt"

func main() {
	input1 := "パトカー"
	input2 := "タクシー"

	runesInput1 := []rune(input1)
	runesInput2 := []rune(input2)
	output := []rune{}

	for i, v := range runesInput1 {
		output = append(output, v, runesInput2[i])
	}

	fmt.Printf("%s", string(output))
}

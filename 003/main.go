package main

import (
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	input := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	re, _ := regexp.Compile("[^\\w ]")
	input = re.ReplaceAllString(input, "")
	inputSplitted := strings.Split(input, " ")
	output := make([]int, 0, len(inputSplitted))

	for _, v := range inputSplitted {
		output = append(output, len(v))
	}

	pp.Print(output)
}

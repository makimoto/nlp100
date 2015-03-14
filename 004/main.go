package main

import (
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	input := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	re, _ := regexp.Compile("[^\\w ]")
	input = re.ReplaceAllString(input, "")
	inputSplitted := strings.Split(input, " ")

	oneCharIndex := make(map[int]bool)
	for _, v := range []int{1, 5, 6, 7, 8, 9, 15, 16, 19} {
		oneCharIndex[v-1] = true
	}

	output := make(map[string]string, 20)

	for i, v := range inputSplitted {
		if oneCharIndex[i] {
			output[v[0:1]] = v
		} else {
			output[v[0:2]] = v
		}
	}

	pp.Print(output)
}

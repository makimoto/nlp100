package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) > 0 {
		for _, fname := range os.Args[1:] {
			fp, _ := os.Open(fname)
			fmt.Printf("%d\t%s\n", countLineCount(fp), fname)
		}
	} else {
		fmt.Printf("%d\t%s\n", countLineCount(os.Stdin), "stdin")
	}
}

func countLineCount(fp *os.File) int {
	scanner := bufio.NewScanner(fp)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
	}
	return lineNum
}

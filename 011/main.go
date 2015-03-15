package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) > 0 {
		for _, fname := range os.Args[1:] {
			fp, _ := os.Open(fname)
			replaceTab(fp)
		}
	} else {
		replaceTab(os.Stdin)
	}
}

func replaceTab(fp *os.File) {
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(strings.Replace(scanner.Text(), "\t", " ", -1))
	}
}

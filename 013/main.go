package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) >= 2 {

		pasteCols(os.Args[1], os.Args[2])

	} else {
		fmt.Fprintf(os.Stderr, "usage: %s col1.txt col2.txt\n", os.Args[0])
	}
}

func pasteCols(col1 string, col2 string) {
	col1Fp, _ := os.Open(col1)
	col2Fp, _ := os.Open(col2)
	defer col1Fp.Close()
	defer col2Fp.Close()

	col1Scanner := bufio.NewScanner(col1Fp)
	col2Scanner := bufio.NewScanner(col2Fp)

	for col1Scanner.Scan() {
		col2Scanner.Scan()
		fmt.Printf("%s\t%s\n", col1Scanner.Text(), col2Scanner.Text())
	}
}

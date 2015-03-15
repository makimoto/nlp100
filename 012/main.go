package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) > 0 {

		fp, _ := os.Open(os.Args[1])
		defer fp.Close()
		cutCols(fp)

	} else {
		cutCols(os.Stdin)
	}
}

func cutCols(fpSrc *os.File) {
	fpCol1, _ := os.OpenFile("col1.txt", os.O_WRONLY|os.O_CREATE, 0600)
	defer fpCol1.Close()
	fpCol2, _ := os.OpenFile("col2.txt", os.O_WRONLY|os.O_CREATE, 0600)
	defer fpCol2.Close()

	srcScanner := bufio.NewScanner(fpSrc)
	col1Writer := bufio.NewWriter(fpCol1)
	defer col1Writer.Flush()
	col2Writer := bufio.NewWriter(fpCol2)
	defer col2Writer.Flush()

	for srcScanner.Scan() {
		cols := strings.Split(srcScanner.Text(), "\t")
		col1Writer.WriteString(fmt.Sprintf("%s\n", cols[0]))
		col2Writer.WriteString(fmt.Sprintf("%s\n", cols[1]))
	}
}

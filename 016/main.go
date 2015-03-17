package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		help()
	} else if len(args) == 1 {
		n, err := strconv.Atoi(os.Args[1])
		errCheck(err)
		split(os.Stdin, "stdin", n)
	} else {
		n, err := strconv.Atoi(os.Args[1])
		errCheck(err)
		fp, err := os.Open(os.Args[2])
		errCheck(err)
		defer fp.Close()
		split(fp, os.Args[2], n)
	}
}

func split(fp *os.File, fname string, n int) {
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	linesPerFile := len(lines) / n
	if linesPerFile*n != len(lines) {
		linesPerFile = len(lines)/(n+1) + 1
	}

	fileNo := 0

	fpWrite, _ := os.OpenFile(fmt.Sprintf("%s-%d.txt", fname, fileNo), os.O_WRONLY|os.O_CREATE, 0600)
	defer fpWrite.Close()

	for i, e := range lines {
		if (linesPerFile*(fileNo+1) <= i) && fileNo != n-1 {
			fileNo++
			fpWrite, _ = os.OpenFile(fmt.Sprintf("%s-%d.txt", fname, fileNo), os.O_WRONLY|os.O_CREATE, 0600)
			defer fpWrite.Close()
		}

		fpWrite.WriteString(fmt.Sprintf("%s\n", e))
	}
}

func help() {
	fmt.Fprintf(os.Stderr, "usage: %s numOfHeadLines [file]\n", os.Args[0])
}

func errCheck(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		help()
		os.Exit(1)
	}
}

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
		tail(os.Stdin, n)
	} else {
		n, err := strconv.Atoi(os.Args[1])
		errCheck(err)
		fp, err := os.Open(os.Args[2])
		errCheck(err)
		defer fp.Close()
		tail(fp, n)
	}
}

func tail(fp *os.File, n int) {
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if n > len(lines) {
		n = len(lines)
	}
	for _, e := range lines[len(lines)-n : len(lines)] {
		fmt.Println(e)
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

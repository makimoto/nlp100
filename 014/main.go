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
		head(os.Stdin, n)
	} else {
		n, err := strconv.Atoi(os.Args[1])
		errCheck(err)
		fp, err := os.Open(os.Args[2])
		errCheck(err)
		defer fp.Close()
		head(fp, n)
	}
}

func head(fp *os.File, n int) {
	scanner := bufio.NewScanner(fp)
	c := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		c++
		if c >= n {
			break
		}
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

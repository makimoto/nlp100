package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	const str = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	fmt.Println(randomStr(str))
}

func randomStr(s string) string {
	words := strings.Fields(s)
	output := make([]string, len(words))
	copy(output, words)

	// slice of index to be randomized
	targetIdxes := make([]int, 0, len(words))

	for i, v := range words {
		if len([]rune(v)) > 4 && i != 0 && i != len(words)-1 {
			targetIdxes = append(targetIdxes, i)
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randList := rand.Perm(len(targetIdxes))

	for i, v := range randList {
		output[targetIdxes[i]] = words[targetIdxes[v]]
	}

	return strings.Join(output, " ")
}

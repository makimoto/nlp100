package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println(cipher("さいきんどう? saikindou?"))
}

func cipher(s string) string {
	runes := []rune(s)
	for i, v := range runes {
		if rune('a') <= v && v <= rune('z') {
			runes[i] = 219 - v
		}
	}
	return string(runes)
}

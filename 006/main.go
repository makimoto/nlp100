package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	bigramsX := charBiGrams("paraparaparadise")
	bigramsY := charBiGrams("paragraph")

	fmt.Println("union:")
	pp.Println(union(bigramsX, bigramsY))

	fmt.Println("intersecton:")
	pp.Println(intersection(bigramsX, bigramsY))

	fmt.Println("difference:")
	pp.Println(difference(bigramsX, bigramsY))

}

func charBiGrams(s string) []string {
	runes := []rune(s)
	var output []string
	for i := 0; i < len(runes)-1; i++ {
		output = append(output, string(runes[i:i+2]))
	}
	return output
}

func union(x []string, y []string) []string {
	xSet := make(map[string]bool)
	for _, v := range x {
		xSet[v] = true
	}

	ySet := make(map[string]bool)
	for _, v := range y {
		ySet[v] = true
	}

	result := make([]string, 0, len(xSet)+len(ySet))

	for k := range xSet {
		if ySet[k] {
			result = append(result, k)
		}
	}
	return result
}

func intersection(x []string, y []string) []string {
	xSet := make(map[string]bool)
	for _, v := range x {
		xSet[v] = true
	}

	ySet := make(map[string]bool)
	for _, v := range y {
		ySet[v] = true
	}

	result := make([]string, 0, len(xSet)+len(ySet))

	for k := range xSet {
		result = append(result, k)
	}

	for k := range ySet {
		if !xSet[k] {
			result = append(result, k)
		}
	}

	return result
}

func difference(x []string, y []string) []string {
	xSet := make(map[string]bool)
	for _, v := range x {
		xSet[v] = true
	}

	ySet := make(map[string]bool)
	for _, v := range y {
		ySet[v] = true
	}

	result := make([]string, 0, len(xSet)+len(ySet))

	for k := range xSet {
		if !ySet[k] {
			result = append(result, k)
		}
	}

	return result
}

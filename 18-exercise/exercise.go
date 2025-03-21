package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

/* REF: https://pkg.go.dev/strings#Fields */
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		if _, ok := m[word]; !ok {
			m[word] = 0
		}
		m[word]++
	} 

	return m
}

func main() {
	wc.Test(WordCount)
}
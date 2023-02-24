package mymaps

import (
	"fmt"
	"strings"
)

func MyTest() {

	results := wordCounts("Hello world")
	fmt.Println(results)
}

func wordCounts(s string) map[string]int {

	words := strings.Fields(s)
	counts := map[string]int{}
	for _, word := range words {
		counts[word]++
	}
	return counts
}

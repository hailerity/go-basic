package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counting_map := make(map[string]int)
	word_list := strings.Fields(s)
	for i := 0; i < len(word_list); i++ {
		counting_map[word_list[i]] += 1
	}

	return couting_map
}

func main() {
	wc.Test(WordCount)
}

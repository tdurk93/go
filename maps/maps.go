package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var wordCounts map[string]int
	wordCounts = make(map[string]int)
	allWords := strings.Fields(s)
	for _, v := range(allWords) {
		wordCounts[v] = wordCounts[v] + 1
	}
	return wordCounts
}

func main() {
	wc.Test(WordCount)
}


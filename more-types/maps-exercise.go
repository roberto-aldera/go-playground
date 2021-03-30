package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	wordcounter := make(map[string]int)
	for _, word := range fields {
		wordcounter[word] += 1
	}
	return wordcounter
}

func run_maps_exercise() {
	wc.Test(WordCount)
}

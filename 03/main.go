package main

import (
	"sort"
	"strings"
)

func main() {
	text := "as ui op op op op ui ui as"
	for _, word := range TopFrequentWords(text) {
		print(word + " ")
	}
}

func TopFrequentWords(text string) [10]string {
	//Count
	words := strings.Split(text, " ")
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++
	}

	//Prepare
	type word struct {
		value string
		count int
	}
	top := make([]word, 0, len(counts))
	for w, count := range counts {
		top = append(top, word{
			value: w,
			count: count,
		})
	}

	//Sort
	sort.Slice(top, func(i, j int) bool {
		if top[i].count == top[j].count {
			return top[i].value > top[j].value
		}
		return top[i].count > top[j].count
	})

	//Format
	length := len(top)
	if length > 10 {
		length = 10
	}

	result := [10]string{}
	for i := 0; i < length; i++ {
		result[i] = top[i].value
	}

	return result
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := strings.ToLower("Hi how are you I am great how about you what are you up to today")
	word_frequencies := make(map[string]int)

	split_sentence := strings.Fields(sentence)

	for _, word := range split_sentence {
		if _, key_exists := word_frequencies[word]; key_exists {
			word_frequencies[word] += 1
		} else {
			word_frequencies[word] = 1
		}
	}

	for word, word_frequency := range word_frequencies {
		fmt.Printf("%s: %d\n", word, word_frequency)
	}
}

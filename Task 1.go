package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Fields(text) // Разбиваем строку на слова
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++ // Увеличиваем счетчик для каждого слова
	}

	return wordCount
}

func main() {
	input := "hello world hello hello everyone"
	result := countWords(input)

	for word, count := range result {
		fmt.Printf("%s: %d\n", word, count)
	}
}
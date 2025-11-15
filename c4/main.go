package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Text processing: count number of words and how many times each word appears in the sentence")
	sentence := "While the basics of regex are well-known, there are advanced techniques and complex patterns that can take your text-processing skills to new heights."

	words := strings.Fields(sentence)
	fmt.Println(words)

	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}

	fmt.Println(count)

}

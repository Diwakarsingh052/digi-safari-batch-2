package main

import (
	"os"
	"strings"
)

type TextAnalysis struct {
	Content       string
	WordFrequency map[string]int
	AvgWordLength float64
	LongestWord   string
}

// AnalyzeText takes a filename as input and reads the content of the file.
// It loops over each word in the `words` slice and performs the following actions:
// - Increments the count for the word in the `frequency` map.
// - Updates the `longest` word if the current word is longer.
// - Adds the length of the current word to `totalLength`.
// After the loop, it calculates the average word length by dividing `totalLength` by the length of the `words` slice.
// Finally, it returns a TextAnalysis struct with the following properties assigned:
// - Content: The original text content of the file.
// - WordFrequency: A map containing the frequency of occurrence of each word.
// - AvgWordLength: The average length of words in the text.
// - LongestWord: The longest word found in the text.

func AnalyzeText(filename string) (TextAnalysis, error) {
	// reading the entire file in memory
	content, err := os.ReadFile(filename)
	if err != nil {
		return TextAnalysis{}, err
	}
	// why is this a flat allocation? Run escape Analysis report

	// s should do the allocation to heap not matter what,
	//but because of function inlining no allocation if the data could be stored on stack
	//s := CreateLargeBuffer()
	//_ = s

	// Convert the byte slice of the file content to a string.
	text := string(content)

	// Split the `text` string into a slice of words.
	words := strings.Fields(text)

	// Initialize a map to hold the frequency of occurrence of each word.
	frequency := make(map[string]int)

	// Initialize totalLength to cumulate the length of all words.
	totalLength := 0

	// Initialize longest to hold the longest word.
	longest := ""

	// Loop over every word in the `words` slice.
	for _, word := range words {
		frequency[word]++ // Increment the count for the word in the frequency map.
		if len(word) > len(longest) {
			longest = word // Update the longest word if the current word is longer.
		}
		totalLength += len(word) // Add the length of the current word to totalLength.
	}

	// Calculating the average word length.
	avgLength := float64(totalLength) / float64(len(words))

	// Return a TextAnalysis instance with properties assigned.
	return TextAnalysis{
		Content:       text,
		WordFrequency: frequency,
		AvgWordLength: avgLength,
		LongestWord:   longest,
	}, nil
}

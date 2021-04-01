package boolTree

import "strings"

type Sentence map[string]bool

func NewSentence(input string) Sentence {
	return createStringMap(input)
}

func createStringMap(input string) Sentence {
	sentence := Sentence{}
	strSlice := strings.Split(input, " ")
	for _, word := range strSlice {
		sentence[word] = true
	}
	return sentence
}

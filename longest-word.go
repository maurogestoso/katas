package katas

import "strings"

// LongestWord returns the length of the longest word in the passed string
func LongestWord(s string) int {
	words := strings.Split(s, " ")
	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return len(longestWord)
}

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	charToWord, wordToChar := make(map[string]string), make(map[string]string)

	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		char := string(pattern[i])

		wordFromMap, ok1 := charToWord[char]
		charFromMap, ok2 := wordToChar[words[i]]

		if !ok1 && !ok2 {
			charToWord[char] = words[i]
			wordToChar[words[i]] = char
		} else if !ok1 || !ok2 {
			return false
		} else if wordFromMap != words[i] || charFromMap != char {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Solution 1 : ", wordPattern("abba", "dog cat cat dog"))
	fmt.Println("Solution 2 : ", wordPattern("abba", "dog cat cat fish"))
	fmt.Println("Solution 3 : ", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println("Solution 4 : ", wordPattern("abba", "dog dog dog dog"))
}

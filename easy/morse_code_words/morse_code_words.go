package main

import "fmt"

var morseCodes = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	transformations := make(map[string]bool)

	for _, word := range words {
		morseCode := ""
		for _, letter := range word {
			morseCode += morseCodes[letter-'a']
		}
		transformations[morseCode] = true
	}
	return len(transformations)
}

func main() {
	fmt.Println("Solution 1 : ", uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

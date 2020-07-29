package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	lengthOfWord := 0
	s = strings.TrimSpace(s)

	i := len(s) - 1
	for i >= 0 && s[i] != ' ' {
		lengthOfWord++
		i--
	}

	return lengthOfWord
}

func main() {
	fmt.Println("Solution 1 : ", lengthOfLastWord("Hello World"))
}

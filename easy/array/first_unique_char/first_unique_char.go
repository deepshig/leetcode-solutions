package main

import "fmt"

func firstUniqChar(s string) int {
	charToIndex := make([]int, 26)
	for _, c := range s {
		charToIndex[c-'a']++
	}

	for i, c := range s {
		if charToIndex[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Solution 1 : ", firstUniqChar("leetcode"))
	fmt.Println("Solution 2 : ", firstUniqChar("loveleetcode"))
}

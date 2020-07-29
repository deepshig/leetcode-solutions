package main

import (
	"fmt"
	"strings"
)

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	placesOfDifference := make([]int, 0, 3)
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			placesOfDifference = append(placesOfDifference, i)
		}
		if len(placesOfDifference) > 2 {
			return false
		}
	}

	if len(placesOfDifference) == 2 {
		first, second := placesOfDifference[0], placesOfDifference[1]
		if A[first] != B[second] || A[second] != B[first] {
			return false
		}
	}

	if len(placesOfDifference) == 1 {
		return false
	}

	if len(placesOfDifference) == 0 {
		for i := 0; i < len(A); i++ {
			if strings.Count(A, string(A[i])) >= 2 {
				return true
			}
		}
		return false
	}

	return true
}

func main() {
	fmt.Println("Solution 1 : ", buddyStrings("ab", "ba"))
	fmt.Println("Solution 2 : ", buddyStrings("ab", "ab"))
	fmt.Println("Solution 3 : ", buddyStrings("aa", "aa"))
	fmt.Println("Solution 4 : ", buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println("Solution 5 : ", buddyStrings("", "aa"))
}

package main

import (
	"fmt"
)

func findPairs(nums []int, k int) int {
	numbersWithOccurrences := make(map[int]int)
	for _, n := range nums {
		if _, ok := numbersWithOccurrences[n]; ok {
			numbersWithOccurrences[n]++
		} else {
			numbersWithOccurrences[n] = 1
		}
	}

	pairs := 0
	for n, o := range numbersWithOccurrences {
		if k == 0 && o > 1 {
			pairs++
		} else if k > 0 {
			partner := n + k
			if _, ok := numbersWithOccurrences[partner]; ok {
				pairs++
			}
		}
	}

	return pairs
}

func main() {
	fmt.Println("Solution 1 : ", findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println("Solution 2 : ", findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println("Solution 3 : ", findPairs([]int{1, 3, 1, 5, 4}, 0))
}

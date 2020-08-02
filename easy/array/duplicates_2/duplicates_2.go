package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	numbersToIndex := make(map[int]int, len(nums))
	for i, n := range nums {
		if prevIndex, ok := numbersToIndex[n]; ok && i-prevIndex <= k {
			return true
		}
		numbersToIndex[n] = i
	}
	return false
}

func main() {
	fmt.Println("Solution 1 : ", containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println("Solution 2 : ", containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println("Solution 3 : ", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

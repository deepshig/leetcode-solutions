package main

import "fmt"

func countPairs(valToNumberOfOccurrences map[int]int) int {
	totalPairs := 0
	for _, occurrences := range valToNumberOfOccurrences {
		allPairs := occurrences * (occurrences - 1)
		increasingOrderPairs := allPairs / 2
		totalPairs = totalPairs + increasingOrderPairs
	}

	return totalPairs
}

func countOccerrences(nums []int) map[int]int {
	valToNumberOfOccurrences := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := valToNumberOfOccurrences[nums[i]]; ok {
			valToNumberOfOccurrences[nums[i]]++
		} else {
			valToNumberOfOccurrences[nums[i]] = 1
		}
	}

	return valToNumberOfOccurrences
}

func numIdenticalPairs(nums []int) int {
	valToNumberOfOccurrences := countOccerrences(nums)
	return countPairs(valToNumberOfOccurrences)
}

func main() {
	fmt.Println("Solution 1 : ", numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println("Solution 2 : ", numIdenticalPairs([]int{1, 1, 1, 1}))
	fmt.Println("Solution 3 : ", numIdenticalPairs([]int{1, 2, 3}))
}

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numberToIndex := make(map[int]int, len(nums))
	for i, n := range nums {
		partner := target - n
		if partnerIndex, ok := numberToIndex[partner]; ok {
			return []int{partnerIndex, i}
		}

		numberToIndex[n] = i
	}
	return nil
}

func main() {
	fmt.Println("Solution 1 : ", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Solution 2 : ", twoSum([]int{3, 3}, 6))
}

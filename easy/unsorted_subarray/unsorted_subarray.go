package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	start, end := 0, len(nums)-1
	needToSort := false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			needToSort = true
			start = i
			break
		}
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			needToSort = true
			end = i
			break
		}
	}

	unsortedSubarray := nums[start : end+1]
	max, min := unsortedSubarray[0], unsortedSubarray[0]
	for i := 1; i < len(unsortedSubarray); i++ {
		if unsortedSubarray[i] > max {
			max = unsortedSubarray[i]
		}
		if unsortedSubarray[i] < min {
			min = unsortedSubarray[i]
		}
	}

	if !needToSort && start == 0 && end == len(nums)-1 {
		return 0
	}

	for start > 0 && nums[start-1] > min {
		start = start - 1
	}

	for end < len(nums)-1 && nums[end+1] < max {
		end = end + 1
	}

	return end - start + 1
}

func main() {
	fmt.Println("Solution 1 : ", findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println("Solution 2 : ", findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println("Solution 3 : ", findUnsortedSubarray([]int{1, 3, 2, 2, 2}))
}

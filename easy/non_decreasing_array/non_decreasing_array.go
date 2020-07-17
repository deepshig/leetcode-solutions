package main

import "fmt"

func checkPossibility(nums []int) bool {
	placeOfViolation := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if placeOfViolation != -1 {
				return false
			}
			placeOfViolation = i
		}
	}

	if placeOfViolation == -1 {
		return true
	}

	indexToChange := placeOfViolation
	if indexToChange == 0 {
		return true
	}

	if nums[indexToChange-1] <= nums[indexToChange+1] {
		return true
	}

	indexToChange = placeOfViolation + 1
	if indexToChange == len(nums)-1 {
		return true
	}

	if (indexToChange < len(nums)-1) && (nums[indexToChange-1] <= nums[indexToChange+1]) {
		return true
	}

	return false
}

func main() {
	fmt.Println("Solution 1 : ", checkPossibility([]int{4, 2, 3}))
	fmt.Println("Solution 2 : ", checkPossibility([]int{4, 2, 1}))
}

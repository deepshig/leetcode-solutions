package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	firstMax, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = nums[i]
		} else if nums[i] > secondMax && nums[i] < firstMax {
			thirdMax = secondMax
			secondMax = nums[i]
		} else if nums[i] > thirdMax && nums[i] < secondMax {
			thirdMax = nums[i]
		}
	}

	if thirdMax > math.MinInt64 {
		return thirdMax
	}

	return firstMax
}

func main() {
	fmt.Println("Solution 1 : ", thirdMax([]int{3, 2, 1}))
	fmt.Println("Solution 2 : ", thirdMax([]int{1, 2}))
	fmt.Println("Solution 3 : ", thirdMax([]int{2, 2, 3, 1}))
}

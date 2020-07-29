package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	firstMax, secondMax := math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > firstMax {
			secondMax, firstMax = firstMax, nums[i]
		} else if nums[i] > secondMax {
			secondMax = nums[i]
		}
	}

	return (firstMax - 1) * (secondMax - 1)
}

func main() {
	fmt.Println("Solution 1 : ", maxProduct([]int{3, 4, 5, 2}))
	fmt.Println("Solution 2 : ", maxProduct([]int{1, 5, 4, 5}))
	fmt.Println("Solution 3 : ", maxProduct([]int{3, 7}))
	fmt.Println("Solution 4 : ", maxProduct([]int{10, 2, 5, 2}))
}

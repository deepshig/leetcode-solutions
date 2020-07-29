package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max_sum := math.MinInt64
	current_sum := 0

	for _, n := range nums {
		current_sum = int(math.Max(float64(n), float64(current_sum+n)))
		max_sum = int(math.Max(float64(max_sum), float64(current_sum)))
	}
	return max_sum
}

func main() {
	fmt.Println("Solution 1 : ", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

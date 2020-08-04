package main

import (
	"fmt"
	"math"
)

func findErrorNums(nums []int) []int {
	duplicate, missing := 0, 0
	for _, n := range nums {
		n = int(math.Abs(float64(n)))
		if nums[n-1] < 0 {
			duplicate = n
		} else {
			nums[n-1] *= -1
		}
	}
	for i, n := range nums {
		if n > 0 {
			missing = i + 1
			break
		}
	}
	return []int{duplicate, missing}
}

func main() {
	fmt.Println("Solution 1 : ", findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println("Solution 1 : ", findErrorNums([]int{2, 2}))
}

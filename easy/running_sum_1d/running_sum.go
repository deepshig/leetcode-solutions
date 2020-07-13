package main

import "fmt"

func runningSum(nums []int) []int {
	length := len(nums)
	sums := make([]int, 0, length)

	sums = append(sums, nums[0])
	for i := 1; i < length; i++ {
		sum := sums[i-1] + nums[i]
		sums = append(sums, sum)
	}

	return sums
}

func main() {
	fmt.Println("Solution 1 : ", runningSum([]int{1, 2, 3, 4}))
	fmt.Println("Solution 2 : ", runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println("Solution 3 : ", runningSum([]int{3, 1, 2, 10, 1}))
}

package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		n := nums[i]
		if nums[n-1] != n {
			nums[i], nums[n-1] = nums[n-1], n
		} else {
			i++
		}
	}

	disappearedNumbers := make([]int, 0, len(nums))
	for i, n := range nums {
		if n != i+1 {
			disappearedNumbers = append(disappearedNumbers, i+1)
		}
	}
	return disappearedNumbers
}

func main() {
	fmt.Println("Solution 1 : ", findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

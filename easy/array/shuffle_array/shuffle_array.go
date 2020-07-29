package main

import "fmt"

func shuffle(nums []int, n int) []int {
	shuffledArray := make([]int, 0, 2*n)

	for i := 0; i < n; i++ {
		shuffledArray = append(shuffledArray, nums[i])
		shuffledArray = append(shuffledArray, nums[i+n])
	}

	return shuffledArray
}

func main() {
	fmt.Println("Solution 1 : ", shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println("Solution 2 : ", shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println("Solution 3 : ", shuffle([]int{1, 1, 2, 2}, 2))
}

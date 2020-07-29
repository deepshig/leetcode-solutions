package main

import "fmt"

func moveZeroes(nums []int) {
	numberOfZeroes := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			numberOfZeroes++
		} else {
			nums[i-numberOfZeroes] = nums[i]
		}
	}

	for numberOfZeroes > 0 {
		nums[len(nums)-numberOfZeroes] = 0
		numberOfZeroes--
	}
	return
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println("Solution 1 : ", nums)

	nums = []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println("Solution 2 : ", nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println("Solution 3 : ", nums)
}

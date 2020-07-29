package main

import (
	"fmt"
)

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func rotate(nums []int, k int) {
	if k == 0 || len(nums) == 1 {
		return
	}
	reverse(nums)

	k %= len(nums)
	firstPart := nums[:k]
	lastPart := nums[k:]

	reverse(firstPart)
	reverse(lastPart)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println("Solution 1 : ", arr)

	arr = []int{-1, -100, 3, 99}
	rotate(arr, 2)
	fmt.Println("Solution 2 : ", arr)

	arr = []int{1, 2}
	rotate(arr, 3)
	fmt.Println("Solution 3 : ", arr)

	arr = []int{1, 2, 3}
	rotate(arr, 2)
	fmt.Println("Solution 4 : ", arr)
}

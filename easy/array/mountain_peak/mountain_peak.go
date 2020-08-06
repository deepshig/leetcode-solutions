package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	start, end := 0, len(A)
	for start < end {
		mid := (start + end) / 2
		if A[mid-1] < A[mid] && A[mid] > A[mid+1] {
			return mid
		} else if A[mid] < A[mid+1] {
			start = mid + 1
		} else if A[mid-1] > A[mid] {
			end = mid
		}
	}
	return -1
}

func main() {
	fmt.Println("Solution 1 : ", peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println("Solution 2 : ", peakIndexInMountainArray([]int{0, 2, 1, 0}))
}

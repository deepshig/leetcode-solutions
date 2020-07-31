package main

import "fmt"

func validMountainArray(A []int) bool {
	if len(A) < 3 || A[0] >= A[1] {
		return false
	}

	peakAcheived := false
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		}
		if !peakAcheived {
			if A[i] < A[i-1] {
				peakAcheived = true
			}
		} else {
			if A[i] > A[i-1] {
				return false
			}
		}
	}

	if !peakAcheived {
		return false
	}
	return true
}

func main() {
	fmt.Println("Solution 1 : ", validMountainArray([]int{2, 1}))
	fmt.Println("Solution 2 : ", validMountainArray([]int{3, 5, 5}))
	fmt.Println("Solution 3 : ", validMountainArray([]int{0, 3, 2, 1}))
}

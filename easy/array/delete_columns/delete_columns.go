package main

import "fmt"

func minDeletionSize(A []string) int {
	columnsToDelete := 0
	for i := 0; i < len(A[0]); i++ {
		for j := 1; j < len(A); j++ {
			if A[j][i] < A[j-1][i] {
				columnsToDelete++
				break
			}
		}
	}
	return columnsToDelete
}

func main() {
	fmt.Println("Solution 1 : ", minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println("Solution 2 : ", minDeletionSize([]string{"a", "b"}))
	fmt.Println("Solution 3 : ", minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}

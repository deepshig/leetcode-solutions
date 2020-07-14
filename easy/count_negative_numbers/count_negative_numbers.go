package main

import "fmt"

func countNegatives(grid [][]int) int {
	columns := len(grid[0])

	totalNegatives := 0
	for _, row := range grid {
		if row[0] < 0 {
			totalNegatives += columns
			continue
		}
		if row[columns-1] >= 0 {
			continue
		}
		mid := columns / 2
		if (mid < columns-1) && (row[mid] >= 0) && (row[mid+1] < 0) {
			negatives := columns - mid - 1
			totalNegatives += negatives
			continue
		}
		if (mid > 1) && (row[mid-1] >= 0) && (row[mid] < 0) {
			negatives := columns - mid
			totalNegatives += negatives
			continue
		}
		for i := columns - 1; i >= 0; i-- {
			if row[i] < 0 {
				totalNegatives++
			} else {
				break
			}
		}

	}

	return totalNegatives
}

func main() {
	fmt.Println("Solution 1 : ", countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	fmt.Println("Solution 2 : ", countNegatives([][]int{{3, 2}, {1, 0}}))
	fmt.Println("Solution 3 : ", countNegatives([][]int{{1, -1}, {-1, -1}}))
	fmt.Println("Solution 4 : ", countNegatives([][]int{{-1}}))
	fmt.Println("Solution 5 : ", countNegatives([][]int{{3, -1, -3, -3, -3}, {2, -2, -3, -3, -3}, {1, -2, -3, -3, -3}, {0, -3, -3, -3, -3}}))
}

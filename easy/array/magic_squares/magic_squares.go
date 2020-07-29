package main

import "fmt"

func checkIfMagicSquare(grid [][]int, startRow, startColumn int) bool {
	digitsOccured := make([]bool, 10, 10)
	for i := 0; i < 10; i++ {
		digitsOccured[i] = false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			current_element := grid[startRow+i][startColumn+j]
			if current_element < 1 || current_element > 9 || digitsOccured[current_element] {
				return false
			}
			digitsOccured[current_element] = true
		}
	}

	firstRowSum := grid[startRow][startColumn] + grid[startRow][startColumn+1] + grid[startRow][startColumn+2]
	secondRowSum := grid[startRow+1][startColumn] + grid[startRow+1][startColumn+1] + grid[startRow+1][startColumn+2]
	thirdRowSum := grid[startRow+2][startColumn] + grid[startRow+2][startColumn+1] + grid[startRow+2][startColumn+2]
	firstColumnSum := grid[startRow][startColumn] + grid[startRow+1][startColumn] + grid[startRow+2][startColumn]
	secondColumnSum := grid[startRow][startColumn+1] + grid[startRow+1][startColumn+1] + grid[startRow+2][startColumn+1]
	thirdColumnSum := grid[startRow][startColumn+2] + grid[startRow+1][startColumn+2] + grid[startRow+2][startColumn+2]
	firstDiagonalSum := grid[startRow][startColumn] + grid[startRow+1][startColumn+1] + grid[startRow+2][startColumn+2]
	secondDiagonalSum := grid[startRow][startColumn+2] + grid[startRow+1][startColumn+1] + grid[startRow+2][startColumn]

	if firstRowSum == secondRowSum && firstRowSum == thirdRowSum && firstRowSum == firstColumnSum && firstRowSum == secondColumnSum && firstRowSum == thirdColumnSum && firstRowSum == firstDiagonalSum && firstRowSum == secondDiagonalSum {
		return true
	}
	return false
}

func numMagicSquaresInside(grid [][]int) int {
	rows, columns := len(grid), len(grid[0])
	magicSquares := 0

	for i := 0; i < rows-2; i++ {
		for j := 0; j < columns-2; j++ {
			if checkIfMagicSquare(grid, i, j) {
				magicSquares++
			}
		}
	}
	return magicSquares
}

func main() {
	grid := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
	fmt.Println("Solution 1 : ", numMagicSquaresInside(grid))
}

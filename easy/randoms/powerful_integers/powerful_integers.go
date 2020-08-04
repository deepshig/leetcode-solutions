package main

import (
	"fmt"
)

func powerfulIntegers(x int, y int, bound int) []int {
	powerfulIntsSet := make(map[int]struct{}, bound)

	for i := 1; i < bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			powerfulIntsSet[i+j] = struct{}{}
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	powerfulIntsList := make([]int, 0, len(powerfulIntsSet))
	for val := range powerfulIntsSet {
		powerfulIntsList = append(powerfulIntsList, val)
	}
	return powerfulIntsList
}

func main() {
	fmt.Println("Solution 1 : ", powerfulIntegers(2, 3, 10))
	fmt.Println("Solution 2 : ", powerfulIntegers(3, 5, 15))
}

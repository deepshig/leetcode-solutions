package main

import "fmt"

func findMax(array []int) int {
	max := array[0]
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := findMax(candies)
	canHaveMaxCandies := make([]bool, 0, len(candies))

	for i := 0; i < len(candies); i++ {
		candiesNeeded := maxCandies - candies[i]
		if candiesNeeded <= extraCandies {
			canHaveMaxCandies = append(canHaveMaxCandies, true)
		} else {
			canHaveMaxCandies = append(canHaveMaxCandies, false)
		}
	}

	return canHaveMaxCandies
}

func main() {
	fmt.Println("Solution 1 : ", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println("Solution 2 : ", kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println("Solution 1 : ", kidsWithCandies([]int{12, 1, 12}, 10))
}

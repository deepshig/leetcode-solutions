package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	firstMin, secondMin := math.MaxInt32, math.MaxInt32
	firstMax, secondMax, thirdMax := math.MinInt32, math.MaxInt32, math.MinInt32

	for _, n := range nums {
		if n <= firstMin {
			firstMin, secondMin = n, firstMin
		} else if n <= secondMin {
			secondMin = n
		}

		if n >= firstMax {
			firstMax, secondMax, thirdMax = n, firstMax, secondMax
		} else if n >= secondMax {
			secondMax, thirdMax = n, secondMax
		} else if n >= thirdMax {
			thirdMax = n
		}
	}

	product1 := firstMin * secondMin * firstMax
	product2 := firstMax * secondMax * thirdMax

	if product1 > product2 {
		return product1
	}
	return product2
}

func main() {
	fmt.Println("Solution 1 : ", maximumProduct([]int{1, 2, 3}))
	fmt.Println("Solution 2 : ", maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println("Solution 3 : ", maximumProduct([]int{-4, -3, -2, -1, 60}))
}

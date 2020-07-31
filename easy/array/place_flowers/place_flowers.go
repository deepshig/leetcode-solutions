package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	totalPlots := len(flowerbed)

	for i := 0; i < totalPlots; i++ {
		left, right := false, false
		if i == 0 || flowerbed[i-1] == 0 {
			left = true
		}
		if i == totalPlots-1 || flowerbed[i+1] == 0 {
			right = true
		}
		if flowerbed[i] == 0 && left && right {
			flowerbed[i] = 1
			n--
		}
		if n <= 0 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Solution 1 : ", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println("Solution 2 : ", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println("Solution 3 : ", canPlaceFlowers([]int{1}, 1))
	fmt.Println("Solution 4 : ", canPlaceFlowers([]int{0}, 1))
	fmt.Println("Solution 5 : ", canPlaceFlowers([]int{0, 1, 0}, 1))
	fmt.Println("Solution 6 : ", canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println("Solution 7 : ", canPlaceFlowers([]int{0, 0, 0, 1}, 1))
}

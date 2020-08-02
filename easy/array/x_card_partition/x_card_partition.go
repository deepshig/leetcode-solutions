package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func hasGroupsSizeX(deck []int) bool {
	numberToOccurence := make(map[int]int)
	for _, c := range deck {
		if _, ok := numberToOccurence[c]; ok {
			numberToOccurence[c]++
		} else {
			numberToOccurence[c] = 1
		}
	}

	x := numberToOccurence[deck[0]]
	if x < 2 {
		return false
	}

	for _, occurence := range numberToOccurence {
		if gcd(x, occurence) < 2 {
			return false
		}
		x = gcd(x, occurence)
	}
	return true
}

func main() {
	fmt.Println("Solution 1 : ", hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println("Solution 2 : ", hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println("Solution 3 : ", hasGroupsSizeX([]int{1}))
	fmt.Println("Solution 4 : ", hasGroupsSizeX([]int{1, 1}))
	fmt.Println("Solution 5 : ", hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
}

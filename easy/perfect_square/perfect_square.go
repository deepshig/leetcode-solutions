package main

import "fmt"

func isPerfectSquare(num int) bool {
	i := 1
	for i*i <= num {
		if i*i == num {
			return true
		}
		i++
	}
	return false
}

func main() {
	fmt.Println("Solution 1 : ", isPerfectSquare(16))
	fmt.Println("Solution 1 : ", isPerfectSquare(14))
}

package main

import "fmt"

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println("Solution 1 : ", canWinNim(4))
	fmt.Println("Solution 2 : ", canWinNim(5))
	fmt.Println("Solution 3 : ", canWinNim(8))
	fmt.Println("Solution 3 : ", canWinNim(16))
}

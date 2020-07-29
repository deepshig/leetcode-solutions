package main

import "fmt"

func balancedStringSplit(s string) int {
	numberOfL := 0
	numberOfR := 0

	balancedSplits := 0
	for _, c := range s {
		if c == 'L' {
			numberOfL++
		} else if c == 'R' {
			numberOfR++
		}
		if numberOfL == numberOfR {
			balancedSplits++
		}
	}
    return balancedSplits
}

func main() {
	fmt.Println("Solution 1 : ", balancedStringSplit("RLRRLLRLRL"))
	fmt.Println("Solution 2 : ", balancedStringSplit("RLLLLRRRLR"))
	fmt.Println("Solution 3 : ", balancedStringSplit("LLLLRRRR"))
	fmt.Println("Solution 4 : ", balancedStringSplit("RLRRRLLRLL"))
}

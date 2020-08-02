package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]struct{})
	for _, j := range J {
		jewels[j] = struct{}{}
	}

	numOfJewels := 0
	for _, s := range S {
		if _, ok := jewels[s]; ok {
			numOfJewels += 1
		}
	}

	return numOfJewels
}

func main() {
	fmt.Println("Solution 1 : ", numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println("Solution 2 : ", numJewelsInStones("z", "ZZ"))
}

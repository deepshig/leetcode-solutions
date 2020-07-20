package main

import "fmt"

func reverseString(s []string) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	a := []string{"h", "e", "l", "l", "o"}
	reverseString(a)
	fmt.Println("Solution 1 : ", a)

	b := []string{"H", "a", "n", "n", "a", "h"}
	reverseString(b)
	fmt.Println("Solution 2 : ", b)
}

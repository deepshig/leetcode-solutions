package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	columnName := ""

	for n > 0 {
		letterNumber := n % 26
		if letterNumber == 0 {
			columnName = "Z" + columnName
			n = n - 26
		} else {
			letter := rune(int('A') + letterNumber - 1)
			columnName = string(letter) + columnName
		}
		n = n / 26
	}

	return columnName
}

func main() {
	fmt.Println("Solution 1 : ", convertToTitle(1))
	fmt.Println("Solution 2 : ", convertToTitle(28))
	fmt.Println("Solution 3 : ", convertToTitle(701))
	fmt.Println("Solution 4 : ", convertToTitle(52))
	fmt.Println("Solution 5 : ", convertToTitle(703))
}

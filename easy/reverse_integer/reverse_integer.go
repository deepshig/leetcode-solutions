package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverseString(str string) string {
	charArr := []rune(str)

	i, j := 0, len(charArr)-1
	for i < j {
		charArr[i], charArr[j] = charArr[j], charArr[i]
		i++
		j--
	}

	return string(charArr)
}

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	negative := false
	if x < 0 {
		negative = true
		x = x * -1
	}

	str := strconv.Itoa(x)
	reverseStr := reverseString(str)
	reverseInt, err := strconv.ParseInt(reverseStr, 10, 32)
	if err != nil {
		return 0
	}

	if reverseInt > math.MaxInt32 || reverseInt < math.MinInt32 {
		return 0
	}
	if negative {
		reverseInt = reverseInt * -1
	}
	return int(reverseInt)
}

func main() {
	fmt.Println("Solution 1 : ", reverse(123))
	fmt.Println("Solution 2 : ", reverse(-123))
	fmt.Println("Solution 3 : ", reverse(120))
	fmt.Println(math.MaxInt32)
}

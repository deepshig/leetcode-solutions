package main

import (
	"fmt"
	"strings"
)

func toLowerCase(str string) string {
	for _, c := range str {
		if (c >= 'A') && (c <= 'Z') {
			str = strings.Replace(str, string(c), string(c+32), 1)
		}
	}
	return str
}

func main() {
	fmt.Println("Solution 1 : ", toLowerCase("Hello"))
	fmt.Println("Solution 2 : ", toLowerCase("here"))
	fmt.Println("Solution 3 : ", toLowerCase("LOVELY"))
}

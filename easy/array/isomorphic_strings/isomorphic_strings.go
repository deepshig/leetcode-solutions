package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sTot, tTos := make(map[rune]rune, len(s)), make(map[rune]rune, len(s))

	for i, sChar := range s {
		tCharFromMap, ok1 := sTot[sChar]
		sCharFromMap, ok2 := tTos[rune(t[i])]
		if !ok1 && !ok2 {
			sTot[sChar] = rune(t[i])
			tTos[rune(t[i])] = sChar
		} else if !ok1 || !ok2 {
			return false
		} else if tCharFromMap != rune(t[i]) || sCharFromMap != sChar {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Solution 1 : ", isIsomorphic("egg", "add"))
	fmt.Println("Solution 2 : ", isIsomorphic("foo", "bar"))
	fmt.Println("Solution 3 : ", isIsomorphic("paper", "title"))
}

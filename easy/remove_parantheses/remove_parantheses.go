package main

import "fmt"

func removeOuterParentheses(S string) string {
	incompleteParanthese := 0
	finalStr := ""
	start := 1

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			incompleteParanthese++
		} else {
			incompleteParanthese--
		}
		if incompleteParanthese == 0 {
			end := i
			finalStr = finalStr + string(S[start:end])
			start = i + 2
		}
	}

	return finalStr
}

func main() {
	fmt.Println("Solution 1 : ", removeOuterParentheses("(()())(())"))
	fmt.Println("Solution 2 : ", removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println("Solution 3 : ", removeOuterParentheses("()()"))
}

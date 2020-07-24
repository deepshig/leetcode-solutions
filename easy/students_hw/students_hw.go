package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	studentsStudying := 0
	for i, e := range endTime {
		if e >= queryTime {
			if startTime[i] <= queryTime {
				studentsStudying++
			}
		}
	}
	return studentsStudying
}

func main() {
	fmt.Println("Solution 1 : ", busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
	fmt.Println("Solution 2 : ", busyStudent([]int{4}, []int{4}, 4))
	fmt.Println("Solution 3 : ", busyStudent([]int{4}, []int{4}, 5))
	fmt.Println("Solution 4 : ", busyStudent([]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7))
	fmt.Println("Solution 5 : ", busyStudent([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5))
}

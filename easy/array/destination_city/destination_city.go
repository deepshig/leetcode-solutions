package main

import "fmt"

func destCity(paths [][]string) string {
	citiesWithOutgoingPaths := make(map[string]bool)
	for _, path := range paths {
		citiesWithOutgoingPaths[path[0]] = true
		if _, ok := citiesWithOutgoingPaths[path[1]]; !ok {
			citiesWithOutgoingPaths[path[1]] = false
		}
	}

	for city, isSource := range citiesWithOutgoingPaths {
		if !isSource {
			return city
		}
	}
	return ""
}

func main() {
	fmt.Println("Solution 1 : ", destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	fmt.Println("Solution 2 : ", destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
	fmt.Println("Solution 3 : ", destCity([][]string{{"A", "Z"}}))
}

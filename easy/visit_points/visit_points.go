package main

import (
	"fmt"
	"math"
)

func areEqual(point1, point2 []int) bool {
	return point1[0] == point2[0] && point1[1] == point2[1]
}

func minTimeToVisitAllPoints(points [][]int) int {
	currentPosition := points[0]
	time := 0
	for _, p := range points {
		if !areEqual(currentPosition, p) {
			deltaX := int(math.Abs(float64(p[0] - currentPosition[0])))
			deltaY := int(math.Abs(float64(p[1] - currentPosition[1])))
			if deltaX > deltaY {
				time += deltaX
			} else {
				time += deltaY
			}
			currentPosition = p
		}
	}
	return time
}

func main() {
	fmt.Println("Solution 1 : ", minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
	fmt.Println("Solution 2 : ", minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))
}

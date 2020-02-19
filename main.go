package main

import (
	"fmt"
	"math"
)

func main() {
	tempValuesList := []int{3, -20, 1, 8, 7, 4, 11, -3}

	t1 := 10
	t2 := -17
	t3 := 6

	nv1 := findNearestTempToVal(tempValuesList, t1)
	nv2 := findNearestTempToVal(tempValuesList, t2)
	nv3 := findNearestTempToVal(tempValuesList, t3)

	fmt.Println(nv1)
	fmt.Println(nv2)
	fmt.Println(nv3)

}

func findNearestTempToVal(tempList []int, val int) int {
	smallestDiff := 1000
	var nearestVal int

	for _, temp := range tempList {
		diff := math.Abs((float64(temp - val)))
		if int(diff) < smallestDiff {
			smallestDiff = int(diff)
			nearestVal = temp
		}

	}
	return nearestVal
}

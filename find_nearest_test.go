package main

import (
	"testing"
)

func TestNearestValueFunction(t *testing.T) {
	tempValuesList := []int{3, -20, 1, 8, 7, 4, 11, -3}

	t1 := 10
	t2 := -17
	t3 := 6
	t4 := 0

	e1 := 11
	e2 := -20
	e3 := 7
	e4 := 1

	a1 := findNearestTempToVal(tempValuesList, t1)
	a2 := findNearestTempToVal(tempValuesList, t2)
	a3 := findNearestTempToVal(tempValuesList, t3)
	a4 := findNearestTempToVal(tempValuesList, t4)

	if a1 != e1 {
		t.Errorf("got %d, expected %d", a1, e1)
	}

	if a2 != e2 {
		t.Errorf("got %d, expected %d", a2, e2)
	}

	if a3 != e3 {
		t.Errorf("got %d, expected %d", a3, e3)
	}

	if a4 == e3 {
		t.Errorf("got %d, expected %d", a4, e4)
	}

}

package main

import (
	"testing"
)

func TestOzon(t *testing.T) {

	a, b := findUnsortedSubarray([]uint{1, 4, 3, 2, 3, 4})
	if !(a == 1 && b == 4) {
		t.Errorf("Test %v, result: %v, %v, expected: 1, 4", 1, a, b)
	} else {
		t.Log("Test 1 Success")
	}

	a, b = findUnsortedSubarray([]uint{2, 6, 4, 8, 10, 9, 15})
	if !(a == 1 && b == 5) {
		t.Errorf("Test %v, result: %v, %v, expected: 1, 5", 2, a, b)
	} else {
		t.Log("Test 2 Success")
	}

	a, b = findUnsortedSubarray([]uint{1, 4, 7, 5, 10, 18, 17, 26, 30, 45, 50, 62})
	if !(a == 2 && b == 6) {
		t.Errorf("Test %v, result: %v, %v, expected: 2, 6", 3, a, b)
	} else {
		t.Log("Test 3 Success")
	}

	a, b = findUnsortedSubarray([]uint{7, 5, 10, 18, 17, 26, 30, 45, 50, 1})
	if !(a == 0 && b == 9) {
		t.Errorf("Test %v, result: %v, %v, expected: 0, 9", 4, a, b)
	} else {
		t.Log("Test 4 Success")
	}

	a, b = findUnsortedSubarray([]uint{1, 4, 7, 5, 10, 18, 17, 26, 30, 45, 50, 2})
	if !(a == 1 && b == 11) {
		t.Errorf("Test %v, result: %v, %v, expected: 1, 11", 5, a, b)
	} else {
		t.Log("Test 5 Success")
	}

	a, b = findUnsortedSubarray([]uint{1, 2, 3, 4, 5, 6, 7})
	if !(a == 0 && b == 0) {
		t.Errorf("Test %v, result: %v, %v, expected: 0, 0", 6, a, b)
	} else {
		t.Log("Test 6 Success")
	}
}

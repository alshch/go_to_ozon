// XVII Задача о поиске не отсортированного подмассива ⭐⭐
// Дана коллекция частично отсортированных целых неотрицательных чисел
// Нужно реализовать алгоритм поиска не отсортированного подмассива
package main

import (
	"sort"
)

func findUnsortedSubarray(array []uint) (left uint, right uint) {
	var isLeftFound, isRightFound bool

	sortArr := make([]uint, len(array))
	copy(sortArr, array)
	sort.Slice(sortArr, func(i, j int) bool { return sortArr[i] < sortArr[j] })
	
	for i, v := range array {
		if v != sortArr[i] && isLeftFound == false {
			left = uint(i)
			isLeftFound = true
		}

		reverseIndex := len(sortArr) - 1 - i
		if array[reverseIndex] != sortArr[reverseIndex] && isRightFound == false {
			right = uint(reverseIndex)
			isRightFound = true
		}

		if isRightFound && isLeftFound {
			break
		}
	}
	return left, right
}

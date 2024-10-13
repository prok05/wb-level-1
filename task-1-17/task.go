// L1.17

// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{4, 8, 10, 2, 9, 7, 3, 5, 6, 1}
	slices.Sort(arr)

	fmt.Println("Index of 5 is: ", binarySearch(arr, 5))
	fmt.Println("Index of 15 is: ", binarySearch(arr, 15))
}

func binarySearch(nums []int, el int) int {
	start, mid, end := 0, 0, len(nums)-1
	for start <= end {
		mid = (start + end) >> 1
		switch {
		case nums[mid] > el:
			end = mid - 1
		case nums[mid] < el:
			start = mid + 1
		default:
			return mid
		}
	}
	return -1
}

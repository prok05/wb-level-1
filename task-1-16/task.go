// L1.16

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{4, 8, 10, 2, 9, 7, 3, 5, 6, 1}
	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left []int
	var right []int
	var mid []int
	log.Println(pivot)

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			mid = append(mid, v)
		}
	}
	return append(append(quickSort(left), mid...), quickSort(right)...)
}

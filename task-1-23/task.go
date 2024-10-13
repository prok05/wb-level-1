// L1.23

// Удалить i-ый элемент из слайса.

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 8

	nums = append(nums[:i], nums[i+1:]...)

	fmt.Println(nums)
}

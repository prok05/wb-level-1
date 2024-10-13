// L1.20

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	fmt.Println(reversePhrase(str))
}

func reversePhrase(str string) string {
	arr := strings.Split(str, " ")

	// что эквивалетно циклу ниже
	//slices.Reverse(arr)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, " ")
}

// L1.8

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func main() {
	num := int64(67)
	fmt.Printf("%064b \n", num)

	// устанавливаем 5 бит в 1
	num = setBit(num, 5, true)
	fmt.Printf("%064b \n", num)
}

func setBit(n int64, i int, bit bool) int64 {
	if bit {
		// устанавливаем бит в 1
		return n | (1 << i)
	} else {
		// устанавливаем бит в 0
		return n &^ (1 << i)
	}
}

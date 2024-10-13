// L1.19

// Разработать программу, которая переворачивает подаваемую на вход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	scanner.Scan()
	str := string(scanner.Text())

	fmt.Println(reverseString(str))
}

func reverseString(str string) string {
	runes := []rune(str)

	// что эквивалетно циклу ниже
	// slices.Reverse(runes)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

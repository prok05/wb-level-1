// L1.26

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
//   abcd — true
//   abCdefAaf — false
//   aabcd — false

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkUniqueSymbols("abcd"))
	fmt.Println(checkUniqueSymbols("abCdefAaf"))
	fmt.Println(checkUniqueSymbols("aabcd"))
	fmt.Println(checkUniqueSymbols("aAbc"))
}

func checkUniqueSymbols(str string) bool {
	// создаем карту для записи значений
	uniques := make(map[rune]bool)

	for _, r := range strings.ToLower(str) {
		// если значение найдено => попался уникальный символ => false
		if uniques[r] {
			return false
		}
		// не найдено => записываем в map
		uniques[r] = true
	}
	return true
}

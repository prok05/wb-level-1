// L1.10

// Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"math"
)

func main() {
	temperature := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// группы для объединения значений
	groups := make(map[int][]float64)

	// проходимся по массиву с температурой
	for _, temp := range temperature {
		// вычисляем группу
		boundary := int(math.Floor(temp/10) * 10)
		// добавляем значение в группу
		groups[boundary] = append(groups[boundary], temp)
	}

	for k, v := range groups {
		fmt.Printf("%d: %v \n", k, v)
	}
}

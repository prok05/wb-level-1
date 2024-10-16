// L1.2

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

package main

import (
	"fmt"
)

func main() {
	// создаем массив
	nums := [5]int{2, 4, 6, 8, 10}

	// создаем небуферизированный канал
	// можно было бы создать буферизированный канал, чтобы запись в канал не блокировалась
	// пока не произойдет чтение в ф-ции main
	results := make(chan int)

	// проходимся по массиву, для каждого значения запускаем горутину
	for _, num := range nums {
		go square(num, results)
	}

	// читаем из канала, выводим значения в консоль
	for range nums {
		fmt.Println(<-results)
	}
}

// ф-ция для рассчета квадрата и записи в канал
func square(num int, results chan int) {
	res := num * num
	results <- res
}

// L1.3

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import "fmt"

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

	// читаем из канала квадраты и складываем их в total
	total := 0
	for range len(nums) {
		total += <-results
	}

	fmt.Println(total)
}

// ф-ция для рассчета квадрата и записи в канал
func square(num int, results chan int) {
	res := num * num
	results <- res
}

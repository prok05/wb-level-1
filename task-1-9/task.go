// L1.9

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}

	numsChan := make(chan int, len(nums))
	dobleNumChan := make(chan int, len(nums))

	var wg sync.WaitGroup

	// пишем в numsChan
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			numsChan <- n
		}(num)
	}

	// читаем из numsChan и пишем в dobleNumChan
	for range len(nums) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num := <-numsChan
			n := num * 2
			dobleNumChan <- n
		}()
	}

	// ожидаем окончания работы горутин
	wg.Wait()

	// закрываем каналы
	close(numsChan)
	close(dobleNumChan)

	// выводим значения из dobleNumChan в stdout
	for res := range dobleNumChan {
		fmt.Println(res)
	}
}

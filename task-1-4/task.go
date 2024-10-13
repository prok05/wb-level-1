// L1.4

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	// проверяем кол-во аргументов командной строки
	if len(os.Args) < 2 {
		log.Fatal("Необходимо указать количество воркеров: go run task.go <количество_воркеров>")
	}

	// считываем кол-во воркеров
	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers <= 0 {
		log.Fatal("Неверное количество воркеров. Укажите положительное число больше 0")
	}

	// создаем канал, в который будем постоянно писать
	results := make(chan int)

	// инициализируем группу ожидания
	var wg sync.WaitGroup

	// для каждого воркера запускаем горутины, которые будут читать из канала
	for w := range workers {
		wg.Add(1)
		go worker(w, results, &wg)
	}

	// создаем канал для сигналов ОС, и подписываемся на сигнал SIGINT
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	// в отдельной горутине запускаем бесконечный цикл, который будет писать рандомное число в канал с интервалом в 1 секунду
	go func() {
		for {
			data := rand.Intn(10)
			results <- data
			time.Sleep(1 * time.Second)
		}
	}()

	// ожидаем сигнала SIGINT
	<-sigs

	// закрываем канал чтобы воркеры завершили чтение
	close(results)

	// ждем окончания работы воркеров
	wg.Wait()

	fmt.Println("Все воркеры закончили работу")
}

// воркер,
func worker(w int, results <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range results {
		fmt.Printf("Получено число: %d, от воркера: %d \n", data, w)
	}
}

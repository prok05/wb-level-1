// L1.6

// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	// 1 Способ - остановка с помощью контекста
	
	ctx, cancel := context.WithCancel(context.Background())
	go withCtxCancel(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(time.Second)
	
	// 2 Способ - остановка записью в канал
	
	//ch := make(chan bool)
	//go withChanSignal(ch)
	//time.Sleep(5 * time.Second)
	//ch <- false

	// 3 Способ - остановка при закрытии канала
	
	//ch := make(chan int)
	//go withChanClosure(ch)
	//time.Sleep(5 * time.Second)
	//close(ch)
	//time.Sleep(time.Second)

	// 4 Способ - остановка с помощью переменной
	
	//var stop bool
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go withVariable(&stop, &wg)
	//time.Sleep(5 * time.Second)
	//stop = true
	//wg.Wait()
}

// Остановка выполнения с помощью контекста
func withCtxCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Завершение горутины с использованием контекста")
			return
		default:
			log.Println("Горутина с использованием контекста")
			time.Sleep(time.Second)
		}
	}
}

// Остановка выполнения с помощью записи в канал
func withChanSignal(ch chan bool) {
	for {
		select {
		case <-ch:
			log.Println("Завершение горутины с помощью записи в канал")
			return
		default:
			log.Println("Горутина с использованием записи в канал")
			time.Sleep(time.Second)
		}
	}
}

// Остановка выполнения с помощью закрытия канала
func withChanClosure(ch chan int) {
	for {
		select {
		case <-ch:
			log.Println("Завершение горутины с помощью закрытия канала")
			return
		default:
			log.Println("Горутина с закрытием канала")
			time.Sleep(time.Second)
		}
	}
}

// Остановка выполнения с помощью переменной
func withVariable(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stop {
			log.Println("Завершение горутины с использованием переменной")
			return
		}
		log.Println("Горутина с использованием переменной")
		time.Sleep(time.Second)
	}
}

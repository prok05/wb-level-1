// L1.18

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// MuCounter счетчик с мьютексом
type MuCounter struct {
	value int
	mu    sync.Mutex
}

// AtomicCounter атомарный счетчик
type AtomicCounter struct {
	value atomic.Int64
}

func main() {
	muCounter := MuCounter{}
	atomicCounter := AtomicCounter{}
	var wg sync.WaitGroup

	// инкрементируем атомарный счетчик
	for range 5 {
		wg.Add(1)
		go func(counter *AtomicCounter, wg *sync.WaitGroup) {
			defer wg.Done()
			counter.value.Add(1)
		}(&atomicCounter, &wg)
	}

	// инкрементируем счетчик с мьютексом
	for range 5 {
		wg.Add(1)
		go func(counter *MuCounter, wg *sync.WaitGroup) {
			defer wg.Done()
			counter.mu.Lock()
			counter.value++
			counter.mu.Unlock()
		}(&muCounter, &wg)
	}

	wg.Wait()
	fmt.Println(muCounter.value)
	fmt.Println(atomicCounter.value.Load())
}

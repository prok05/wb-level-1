// L1.7

// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1 Метод - с использованием sync.RWMutex
	writeWithLock()

	// 2 Метод - с использованием sync.Map
	writeWithSyncMap()
}

func writeWithLock() {
	m := make(map[string]string)
	var mu sync.RWMutex

	for range 15 {
		go func() {
			mu.Lock()
			m["name"] = "oleg"
			mu.Unlock()
		}()
	}

	for range 15 {
		go func() {
			mu.Lock()
			m["name"] = "petr"
			mu.Unlock()
		}()
	}

	for range 15 {
		go func() {
			mu.Lock()
			m["name"] = "ivan"
			mu.Unlock()
		}()
	}

	for range 15 {
		go func() {
			mu.Lock()
			m["name"] = "alex"
			mu.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)

	mu.RLock()
	fmt.Println(m["name"])
	mu.RUnlock()
}

func writeWithSyncMap() {
	sm := sync.Map{}

	for range 15 {
		go func() {
			sm.Store("name", "oleg")
		}()
	}

	for range 15 {
		go func() {
			sm.Store("name", "petr")
		}()
	}

	for range 15 {
		go func() {
			sm.Store("name", "ivan")
		}()
	}

	for range 15 {
		go func() {
			sm.Store("name", "alex")
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println(sm.Load("name"))
}

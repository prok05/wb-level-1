// L1.25

// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Засыпаем...")
	sleep(2 * time.Second)
	fmt.Println("Проснулись!")
}

func sleep(d time.Duration) {
	// ждем получения значения из канала спустя d
	<-time.After(d)
}

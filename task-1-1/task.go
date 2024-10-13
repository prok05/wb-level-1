// L1.1

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type human struct {
	name string
	age  int
}

// Используем композицию - встраиваем структуру Human в структуру Action
type action struct {
	human
}

// Метод speak для структуры Human
func (h *human) speak() {
	fmt.Printf("%s speaks \n", h.name)
}

// Метод printAge для структуры Human
func (h *human) printAge() {
	fmt.Printf("%s is %d \n", h.name, h.age)
}

func main() {
	h := human{name: "Ivan", age: 25}
	action := action{human: h}

	action.speak()
	action.printAge()
}

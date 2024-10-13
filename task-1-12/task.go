// L1.12

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
	"strings"
)

// Set структура множества
type Set[T comparable] struct {
	elems map[T]struct{}
}

// New конструктор для Set
func New[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.elems = make(map[T]struct{})
	return s
}

// Add для добавления элемента
func (s *Set[T]) Add(v T) {
	s.elems[v] = struct{}{}
}

// String для удобочитаемого вывода
func (s *Set[T]) String() string {
	var elems []string
	for v := range s.elems {
		elems = append(elems, fmt.Sprintf("%v", v))
	}
	return "{" + strings.Join(elems, ", ") + "}"
}

func main() {
	// массив строк
	strArr := [5]string{"cat", "cat", "dog", "cat", "tree"}

	// создание множества
	strSet := New[string]()

	// добавляем элементы во множество
	for _, v := range strArr {
		strSet.Add(v)
	}

	fmt.Println(strSet)
}

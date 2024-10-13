// L1.11

// Реализовать пересечение двух неупорядоченных множеств.

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

// Has для проверка наличия элемента
func (s *Set[T]) Has(v T) bool {
	_, ok := s.elems[v]
	return ok
}

// Intersect для получения пересечения
func (s *Set[T]) Intersect(newS *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.elems {
		if !newS.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
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
	// создаем первое множество и добавляем элементы
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	fmt.Println(set1)

	// создаем второе множество и добавляем элементы
	set2 := New[int]()
	set2.Add(2)
	set2.Add(3)
	set2.Add(4)
	fmt.Println(set2)

	// получаем пересечение первого и второго множеств
	interSet := set1.Intersect(set2)
	fmt.Println(interSet)
}

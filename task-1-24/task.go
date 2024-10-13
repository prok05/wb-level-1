// L1.24

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(1, 2)

	dist := calcDistance(point1, point2)
	fmt.Println("Расстояние между point1 и point2:", dist)
}

func calcDistance(p1, p2 *Point) float64 {
	// вычисляем расстояние по формуле AB = √(xb - xa)2 + (yb - ya)2
	return math.Sqrt(math.Pow(p2.GetX()-p1.GetX(), 2) + math.Pow(p2.GetY()-p1.GetY(), 2))
}

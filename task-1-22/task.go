// L1.22

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 26)
	b := big.NewInt(1 << 24)

	addition := new(big.Int)
	subtraction := new(big.Int)
	division := new(big.Int)
	multiplication := new(big.Int)

	addition.Add(a, b)
	subtraction.Sub(a, b)
	division.Div(a, b)
	multiplication.Mul(a, b)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	fmt.Println("Результат a + b:", addition)
	fmt.Println("Результат a - b:", subtraction)
	fmt.Println("Результат a / b:", division)
	fmt.Println("Результат a * b:", multiplication)
}

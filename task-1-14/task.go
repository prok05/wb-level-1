// L1.14

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intVar interface{}
	var stringVar interface{}
	var boolVar interface{}
	var chanVar interface{}

	intVar = 5
	stringVar = "hello"
	boolVar = true
	chanVar = make(chan int)

	fmt.Println("Type of intVar:", reflect.TypeOf(intVar))
	fmt.Println("Type of stringVar:", reflect.TypeOf(stringVar))
	fmt.Println("Type of boolVar:", reflect.TypeOf(boolVar))
	fmt.Println("Type of chanVar:", reflect.TypeOf(chanVar))

	fmt.Println()

	fmt.Println("Type of intVar:", reflect.ValueOf(intVar).Kind())
	fmt.Println("Type of stringVar:", reflect.ValueOf(stringVar).Kind())
	fmt.Println("Type of boolVar:", reflect.ValueOf(boolVar).Kind())
	fmt.Println("Type of chanVar:", reflect.ValueOf(chanVar).Kind())

	fmt.Println()

	fmt.Printf("Type of intVar: %T \n", intVar)
	fmt.Printf("Type of stringVar: %T \n", stringVar)
	fmt.Printf("Type of boolVar: %T \n", boolVar)
	fmt.Printf("Type of chanVar: %T \n", chanVar)
}

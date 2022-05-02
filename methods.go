package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) MethodWithReturn() int {
	return len(m)
}

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	value.MethodWithParameters(4, true)
	fmt.Println(value.MethodWithReturn())
	anotherValue := MyType("another value")
	anotherValue.sayHi()
	anotherValue.MethodWithParameters(-17, false)
	fmt.Println(anotherValue.MethodWithReturn())
	
	fmt.Println()
	
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}
package main

import "fmt"

type long int

func (a long) Add02(b long) (result long) {
	result = a + b
	return
}

func main() {

	result1 := Add(1, 2)
	fmt.Println("result1 =", result1)

	var a long = 2
	result2 := a.Add02(3)
	fmt.Println("result2 =", result2)
}

func Add(a, b int) (result int) {
	result = a + b
	return
}

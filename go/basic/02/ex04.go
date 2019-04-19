package main

import "fmt"

type FuncType func(int, int) int

func add(a, b int) (result int) {
	result = a + b
	return
}

func sub(a, b int) (result int) {
	result = a - b
	return
}

func clac(a, b int, fTest FuncType) (result int) {
	result = fTest(a, b)
	return
}

func main() {
	result1 := clac(1, 2, add)
	result2 := clac(2, 1, sub)
	fmt.Println(result1)
	fmt.Println(result2)
}

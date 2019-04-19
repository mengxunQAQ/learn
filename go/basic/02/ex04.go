/* 回调函数和多态 */
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
	result = fTest(a, b) // callback function
	return
}

func main() {
	result1 := clac(1, 2, add) // 多态
	result2 := clac(2, 1, sub) // 多态
	fmt.Println(result1)
	fmt.Println(result2)
}

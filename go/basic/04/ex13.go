/*
方法表达式
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintInfo() {
	fmt.Println(p)
}

func main() {
	p := Person{"Bob", 22}

	pFunc := Person.PrintInfo // 方法表达式
	pFunc(p)                  // 和方法值不同，调用时需要显式的传递接受者
}

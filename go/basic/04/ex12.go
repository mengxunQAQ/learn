/*
方法值：提前保存方法，方便后面调用
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

	p.PrintInfo()

	pFunc := p.PrintInfo // pFunc 就是方法值，提前保存接受者，后面直接调用方法值即可
	pFunc()
}

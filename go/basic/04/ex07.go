/* 结构体和方法 */
package main

import "fmt"

type Person struct {
	name   string
	age    int
	height int
}

func (tmp Person) PrintInfo() {
	fmt.Println("tmp =", tmp)
}

func (p *Person) SetInfo(name string, age, height int) {
	p.name = name
	p.age = age
	p.height = height
}

func main() {
	// p就是对象，结构体的变量就是对象的成员变量，方法就是对象的方法
	var p Person
	// p类似于Python里的self
	(&p).SetInfo("Bob", 22, 190)
	p.PrintInfo()
}

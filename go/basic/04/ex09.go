/*
方法集：一个类型可以使用的方法的集合
* 指针类型 和 普通类型之间，编译器会自动转换
* 比如 p1是普通类型，但是可以调用SetInfoPointer
* p2是指针类型，同样也可以调用SetInfoValue
*/
package main

import "fmt"

type Person struct {
	name string
	id   int
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoValue")
}

func main() {
	p1 := Person{name: "River"}
	p1.SetInfoValue()
	p1.SetInfoPointer() // 先转换p1为&p1，再调用方法

	p2 := &Person{name: "Bob"}
	p2.SetInfoValue()
	p2.SetInfoPointer() // 先转换p2为*p2，再调用方法
}

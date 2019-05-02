/*
方法的重写
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	id    int
	score int
}

func (p Person) PrintInfo() {
	fmt.Println(p)
}

func (s Student) PrintInfo() {
	fmt.Println(s)
}

func main() {
	p := Person{name: "Bob"}
	p.PrintInfo()

	s := Student{p, 1024, 99}
	s.PrintInfo()        // 调用重写的Student的方法
	s.Person.PrintInfo() // 要调Person的，就必须显式声明
}

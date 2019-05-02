/*
 方法的继承：通过匿名字段，同时会把方法继承过来
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	score int
	id    int
}

func (p *Person) PrintInfo() {
	fmt.Printf("name=%s, age=%d\n", p.name, p.age)
}

func main() {
	p := Person{"Bob", 22}
	p.PrintInfo()

	s := Student{p, 99, 10010}
	s.PrintInfo()
}

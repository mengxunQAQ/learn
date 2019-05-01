/* 指针类型匿名字段 */
package main

import "fmt"

type Person struct {
	name   string
	age    int
	height int
}

type Student struct {
	*Person // 指针类型匿名字段
	int     // 也是匿名字段，只不过是基础数据类型的匿名，非结构体匿名
	score   int
	name    string // 和 Person 同名
}

func main() {
	s1 := Student{&Person{name: "mike"}, 555, 100, "Jack"}
	fmt.Printf("s1 = %+v\n", s1)
	fmt.Printf("s1.Person = %+v\n", *s1.Person)

	var s2 Student
	s2.Person = new(Person)
	s2.Person.name = "Hack"
	fmt.Printf("s2 = %+v\n", s2)
}

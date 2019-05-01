/* 成员变量的操作 */
package main

import "fmt"

type Person struct {
	name   string
	age    int
	height int
}

type Student struct {
	Person // 只有类型，匿名字段
	id     int
	score  int
}

func main() {
	s1 := Student{Person{"Eric", 19, 180}, 101, 98}
	s1.name = "Bob"
	s1.age = 22
	s1.height = 190
	s1.id = 102
	s1.score = 99
	fmt.Printf("s1 = %+v\n", s1) // +v 详细显示

	s1.Person = Person{"go", 13, 188}
	fmt.Printf("s1 = %+v\n", s1) // +v 详细显示

}

/* 同名字段的覆盖 */
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
	name   string // 和 Person 同名
}

func main() {
	var s Student
	s.name = "mike" // 就近原则，会优先赋予Student的name
	s.age = 19
	s.id = 100

	fmt.Printf("s = %+v\n", s)

	s.Person.name = "Jack"
	fmt.Printf("s = %+v\n", s)
}

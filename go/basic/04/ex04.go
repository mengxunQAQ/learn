/* 非结构体的匿名 */
package main

import "fmt"

type Person struct {
	name   string
	age    int
	height int
}

type Student struct {
	Person // 只有类型，匿名字段
	int    // 也是匿名字段，只不过是基础数据类型的匿名，非结构体匿名
	score  int
	name   string // 和 Person 同名
}

func main() {
	s := Student{Person{name: "mike"}, 555, 100, "Jack"}

	fmt.Printf("s = %+v\n", s)
	s.int = 666 // 没有变量名，类型名就是变量名
	fmt.Printf("s = %+v\n", s)
}

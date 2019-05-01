/* 匿名字段初始化 */
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
	s1 := Person{name: "mike", height: 188}
	fmt.Println("s1.height =", s1.height)

	s2 := Student{Person{"Eric", 19, 180}, 101, 98}
	fmt.Printf("s2 = %+v\n", s2) // +v 详细显示

	s3 := Student{s1, 102, 99}
	fmt.Printf("s3 = %+v\n", s3)
}

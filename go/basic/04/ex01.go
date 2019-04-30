/* 匿名字段初始化 */
package main

import "fmt"

type Person struct {
	name   string
	age    int
	height int
}

func main() {
	s1 := Person{name: "mike", height: 188}
	fmt.Println("s1.height =", s1.height)
}

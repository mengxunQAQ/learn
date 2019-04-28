/* 结构体初始化 */
package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// (1 顺序初始化，每个成员都要初始化
	s1 := Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println("s1 =", s1)
	// (2 指定成员初始化, 未初始化的自动赋值为0
	s2 := Student{name: "mike", addr: "shanghai"}
	fmt.Println("s2 =", s2)
}

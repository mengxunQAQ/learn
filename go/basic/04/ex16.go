/*
接口的继承
*/
package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Person interface {
	Humaner
	sing()
}

type Student struct {
	name string
	age  int
}

func (s Student) sayhi() {
	fmt.Println("hi")
}

func (s Student) sing() {
	fmt.Println("sing")
}

func main() {
	var i Person // 定义一个Person接口类型的变量
	s := Student{"Bob", 22}
	i = s
	i.sayhi() // 继承过来的方法
	i.sing()
}

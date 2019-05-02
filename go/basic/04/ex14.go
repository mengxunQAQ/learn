/*
接口：一系列方法的集合
*/
package main

import "fmt"

// 接口
type Humaner interface {
	sayhi()
}

// 1
type Teacher struct {
	name string
	age  int
}

func (t Teacher) sayhi() {
	fmt.Printf("Teacher:%s\n", t.name)
}

// 2
type Student struct {
	score int
	id    int
}

func (s Student) sayhi() {
	fmt.Printf("Student:%d\n", s.score)
}

// 3
type Mystr string

func (s Mystr) sayhi() {
	fmt.Printf("Mystr:%s\n", s)
}

func main() {
	var i Humaner

	s := Student{99, 10010}
	i = s
	i.sayhi()

	t := Teacher{"Alice", 31}
	i = t
	i.sayhi()

	var str Mystr = "Hi"
	i = str
	i.sayhi()
}

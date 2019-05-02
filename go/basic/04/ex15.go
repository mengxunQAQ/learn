/*
多态
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

func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := Student{99, 10010}
	t := Teacher{"Bob", 40}
	var str Mystr = "Hi"

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(str)

	slice := make([]Humaner, 3)
	slice[0] = s
	slice[1] = t
	slice[2] = str

	for _, i := range slice {
		WhoSayHi(i)
	}
}

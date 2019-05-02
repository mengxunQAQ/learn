/* 所有的参数传递都是值传递 */
package main

import "fmt"

type Person struct {
	name string
	id   int
	age  int
}

func (p2 *Person) test() {
	fmt.Printf("p2 = %p\n", &p2)
}

func main() {
	s1 := Person{name: "go"}
	p := &s1
	p.test()
	fmt.Printf("p = %p\n", &p)
}

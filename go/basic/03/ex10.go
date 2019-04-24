/* 数组作参数传递引用 */
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify(&a)
	fmt.Println("main: a=", a)
}

func modify(p *[5]int) {
	(*p)[0] = 555
	fmt.Println("modify: a=", *p)
}

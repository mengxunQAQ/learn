/* 变量 常量的各种声明方式 */
package main

import "fmt"

func main() {
	var a int = 1 //  初始化

	var b int //  声明
	b = 2     // 赋值

	c := 3 // 自动推导

	var (
		d int // 多个声明
		e int
	)
	d, e = 4, 5 // 多个同时赋值

	var (
		f = 5 // 自动推导，不用:=
		//i int
	)

	//const i
	//i = 2
	//const i = 2  const必须初始化，不能之后再赋值

	const (
		g     = 6 // 自动推导，不用:=
		h int = 7
	)

	fmt.Printf("a=%d b=%d c=%d d=%d e=%d f=%d g=%d h=%d\n", a, b, c, d, e, f, g, h)
}

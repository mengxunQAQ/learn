/* 闭包 */
package main

import "fmt"

func test02() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func test01() int {
	var x int // 没初始化，值为0
	x++
	return x * x
}

func main01() {
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
}

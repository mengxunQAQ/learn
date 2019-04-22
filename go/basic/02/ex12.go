/* init */
package main

import "fmt"

func init() {
	fmt.Println("init...")
}

func main() {
	a := 1
	b := 1

	func(x int) {
		x = 2 // 传值
		b = 2 // 传引用
	}(a)

	fmt.Println("a =", a) // 1
	fmt.Println("b =", b) // 2
}

/* 数组的比较和赋值 */
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3}
	b := [5]int{1, 2, 2: 3}
	c := [5]int{1, 2}

	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c) // 相同类型的数组可以进行 == != 这两种比较，> < 这些则不可以

	var d [5]int
	d = a // 同类型的数组可以赋值
	fmt.Println("d =", d)
}

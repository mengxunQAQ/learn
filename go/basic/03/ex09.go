/* 数组当参数 */
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} // 初始化

	modify(a) // 数组当参数

	fmt.Println("main: a=", a)
}

func modify(a [5]int) {
	a[0] = 555
	fmt.Println("modify: a=", a)
}

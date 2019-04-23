/* 数组初始化 */
package main

import "fmt"

func main() {
	// 初始化 = 声明 + 赋值
	// 1.全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a =", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b =", b)

	// 2.部分初始化，未初始化的部分为0
	c := [5]int{1, 2, 3}
	fmt.Println("c =", c)

	// 3.指定特定元素初始化
	d := [5]int{2: 10, 4: 2}
	fmt.Println("d =", d)
}

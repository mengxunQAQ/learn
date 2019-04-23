/* 数组的基本操作 */
package main

import "fmt"

func main() {
	var a [10]int // 定义数组元素个数必须是*常量*

	fmt.Println("len(a) =", len(a))

	a[0] = 1
	i := 1
	a[i] = 2 // 下标可以是变量或者常量

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
		fmt.Println(a[i])
	}
}

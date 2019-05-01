/* method */
package main

import "fmt"

// 方法：给某个类型绑定一个函数
type long int

// a是接受者，接受者也是参数
func (a long) Add02(b long) (result long) {
	result = a + b
	return
}

func main() {

	result1 := Add(1, 2)
	fmt.Println("result1 =", result1)

	var a long = 2
	result2 := a.Add02(3)
	fmt.Println("result2 =", result2)
}

func Add(a, b int) (result int) {
	result = a + b
	return
}

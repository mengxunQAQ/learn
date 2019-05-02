/*
空接口：所有类型都实现里空接口，即意味着所有类型都可以赋值给空接口
应用：func test(args ...interface{}) ==> 不定参数个数，不定类型
*/
package main

import "fmt"

func main() {
	var s interface{} = 1
	fmt.Println(s)

	s = "abc"
	fmt.Println(s)
}

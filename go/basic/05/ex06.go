/* 原生字符串 */
package main

import "fmt"

func main() {
	a := `aaa\n` // 不会进行转义，写成什么样就是什么样
	b := "aaa\n"

	fmt.Println(a)
	fmt.Println(b)
}

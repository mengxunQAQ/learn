/* defer */
package main

import "fmt"

func main() {
	fmt.Println(1)
	defer fmt.Println(2) // 多个defer先声明的后调用
	defer fmt.Println(3) // 无论发生何种错误，defer总是能被调用
	fmt.Println(4)
}

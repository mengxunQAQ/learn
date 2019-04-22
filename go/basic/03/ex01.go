/* 指针 */
package main

import "fmt"

func main() {
	var a int = 10              // 变量有两个属性，1.本身的数据 2.数据的地址
	fmt.Printf("a = %d\n", a)   // 数据
	fmt.Printf("&a = %p\n", &a) // 地址

	var p *int
	p = &a                    //指针
	fmt.Printf("p = %p\n", p) // 指针

	*p = 666 // *p 指针指向的值
	fmt.Printf("*p = %d\n", *p)
}

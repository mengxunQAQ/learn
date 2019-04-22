/* 局部变量和全局变量 */
package main

import "fmt"

// a := 10 短变量声明不可以在函数外
// 全局变量在任何地方都能使用
var a int

func test() {
	fmt.Println("test: a=", a)
}

func main() {
	// 定义在{}里面的变量就是局部变量
	// 执行到声明变量才开始分配空间，离开作用域自动释放
	// 作用域，变量其作用的范围
	a = 10
	fmt.Println("main: a=", a)
	test()

	{
		i := 1
		fmt.Println(i)
	}
	// i = 111 不可以

	if flag := 3; flag == 3 {
		fmt.Println("flog =", flag)
	} // flag只属于if语句
	//flag = 4
}

/*函数以及参数*/
package main

import "fmt"

func main() {
	MyFunc01(100, 1, 2, 3)
}

/*
不定参数必须在参数列表的最后
*/
func MyFunc01(a int, args ...int) {
	fmt.Println("a =", a)
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
	//  传递全部参数
	//	MyFunc03(args...)

	//	传递部分参数
	MyFunc03(args[:2]...)
}

func MyFunc02(a, b int) {
}

func MyFunc03(tmp ...int) {
	for i, data := range tmp {
		fmt.Printf("tmp[%d] = %d\n", i, data)
	}
}

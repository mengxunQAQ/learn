/* Goexit 退出当前协程序· */
package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccc")
	runtime.Goexit()
	fmt.Println("ddd")
}

func main() {
	go func() {
		fmt.Println("aaa")
		test()
		fmt.Println("bbb")
	}()

	// 死循环，不让主协程退出
	for {
	}
}

/* 协程的同步 */
package main

import (
	"fmt"
	"time"
)

//共享资源
func Printer(s string) {
	for _, char := range s {
		fmt.Printf("%c", char)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello:)")
}

func person2() {
	Printer("world")
}

func main() {
	//person1()
	//person2()
	go person1()
	go person2()

	// 不让主协程退出
	for {
	}
}

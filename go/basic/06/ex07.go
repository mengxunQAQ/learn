/* 通过channel实现同步 */
package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

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
	ch <- 555
}

func person2() {
	<-ch
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

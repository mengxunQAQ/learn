/* 无缓冲区channel */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("new:", i)
			ch <- i // 无缓冲区channel，如果存的数据没有协程来取，则会被阻塞
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		<-ch // 如果无数据会被阻塞在这里
		fmt.Println("main:", i)
	}
}

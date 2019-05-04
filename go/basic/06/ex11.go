/* 有缓冲区的channel */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // 有缓冲区channel，大小为3

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("new:", i)
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println("main:", i)
		<-ch
	}

}

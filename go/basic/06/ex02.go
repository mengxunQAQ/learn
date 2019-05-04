/* 主协程退出导致子协程退出 */
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("new:", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main:", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}

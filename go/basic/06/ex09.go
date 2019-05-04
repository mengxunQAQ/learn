/* 通过channel进行数据交互 */
package main

import "fmt"

var ch = make(chan string)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
		}

		ch <- "new finish"
	}()

	str := <-ch
	fmt.Println(str)
}

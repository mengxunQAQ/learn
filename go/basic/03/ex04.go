/* 数组 */
package main

import "fmt"

func main() {
	var id [50]int // go里面未初始化的变量，值默认是0

	for i := range id {
		id[i] = i + 1
		fmt.Println(id[i])
	}
}

/* 命令行参数 */
package main

import "fmt"
import "os"

func main() {
	// 接受用户传递的参数，参数都是以字符串形式传递
	list := os.Args

	n := len(list)
	fmt.Println("n =", n)

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}

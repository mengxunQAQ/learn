/* 删除map里的元素 */
package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go"}
	fmt.Println("m =", m)

	delete(m, 1) // 删除map里键为1的元素
	fmt.Println("m =", m)
}

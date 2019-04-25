/* 切片和数组的区别 */
package main

import "fmt"

func main() {
	a := [5]int{} // 数组不能修改长度
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	s := []int{} // 切片的长度或容量不固定
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s = append(s, 11) // 给切片末尾追加一个成员
	fmt.Printf("append: len = %d, cap = %d\n", len(s), cap(s))
}

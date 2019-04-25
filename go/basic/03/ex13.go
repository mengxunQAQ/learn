/* 切片的创建 */
package main

import "fmt"

func main() {
	// 1-自动推导创建切片
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("s1 =", s1)

	// 2-make函数创建切片，make（切片类型， 长度， 容量）
	s2 := make([]int, 5, 10)
	fmt.Println("len(s2) = %d, cap(s2) = %d\n", len(s2), cap(s2))

	// 3-可以不带 容量， 此时 容量 = 长度
	s3 := make([]int, 5)
	fmt.Println("len(s3) = %d, cap(s3) = %d\n", len(s3), cap(s3))
}

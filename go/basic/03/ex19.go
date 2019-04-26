/* map的基本操作 */
package main

import "fmt"

func main() {
	// 定义变量，类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 =", m2)
	fmt.Println("m1 =", len(m1))

	// 可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 =", m2)
	fmt.Println("len =", len(m2))

	// 通过make创建时，可以提前分配好空间，这样就不用重复分配空间以省去迁移数据的时间
	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	fmt.Println("m3 =", m3)
	fmt.Println("len =", len(m3))

	// 初始化
	m4 := map[int]string{1: "mike", 2: "go"}
	fmt.Println("m4 =", m4)
	fmt.Println("len =", len(m4))
}

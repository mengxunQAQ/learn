/* 切片copy */
package main

import "fmt"

func main() {
	src := []int{1, 2}
	dst := []int{6, 6, 6, 6, 6}

	copy(dst, src)            // src -> dst
	fmt.Println("dst =", dst) // [1, 2, 6, 6, 6]
}

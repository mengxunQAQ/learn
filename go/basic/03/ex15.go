/* 底层数组与切片的关系 */
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := a[2:5]
	s1[1] = 555
	fmt.Println("a =", a)

	s2 := s1[2:7]
	s2[2] = 666
	fmt.Println("a =", a)
}

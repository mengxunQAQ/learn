/* 切片append函数 */
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:2]

	s2 = append(s2, 100)
	fmt.Println("s1 =", s1)
}

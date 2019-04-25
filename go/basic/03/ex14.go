/* 切片截取 */
package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := array[:3]
	fmt.Printf("cap(s1) = %d\n", cap(s1)) //10
	fmt.Printf("len(s1) = %d\n", len(s1)) //3

	s2 := array[3:]
	fmt.Printf("cap(s2) = %d\n", cap(s2)) //7
	fmt.Printf("len(s2) = %d\n", len(s2)) //7

	s3 := array[1:4:5]
	fmt.Printf("cap(s3) = %d\n", cap(s3)) //4
	fmt.Printf("len(s3) = %d\n", len(s3)) //3
}

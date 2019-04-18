/*有参数有返回值的函数*/
package main

import "fmt"

func main() {
	max, min := MaxAndMin(1, 2)
	fmt.Printf("max = %d, min = %d\n", max, min)
}

func MaxAndMin(x int, y int) (max int, min int) {
	if x > y {
		max = x
		min = y
	} else {
		min = x
		max = y
	}

	return
}

/* defer 和 匿名函数 */
package main

import "fmt"

func main() {
	a := 1
	b := 2

	defer func(a, b int) {
		fmt.Printf("defer1: a = %d, b = %d\n", a, b)
	}(a, b) // 自调

	defer func() {
		fmt.Printf("defer2: a = %d, b = %d\n", a, b)
	}()

	a = 10
	b = 20

	fmt.Printf("external: a = %d, b = %d\n", a, b)
}

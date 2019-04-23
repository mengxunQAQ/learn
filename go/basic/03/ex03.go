/* 传值和传指针 */
package main

import "fmt"

func main() {
	a, b := 10, 20

	swap(&a, b)
	fmt.Println(a, b)
}

func swap(a *int, b int) {
	*a, b = b, *a
}

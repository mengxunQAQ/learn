/* recover */
package main

import "fmt"

func test(x int) {
	var a [10]int
	a[x] = 1
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	test(20)
}

/* 枚举 */
package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
	)
	fmt.Printf("a = %d b = %d\n", a, b)

	const c = iota
	fmt.Printf("c = %d\n", c)

	const (
		d = iota
		e
		f
	)
	fmt.Printf("d = %d e = %d f = %d\n", d, e, f)

	const (
		j          = iota
		h1, h2, h3 = iota, iota, iota
		i          = iota
	)
	fmt.Printf("j = %d h1 = %d h2 = %d h3 = %d i = %d\n", j, h1, h2, h3, i)

}

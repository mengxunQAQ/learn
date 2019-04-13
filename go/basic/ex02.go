/* 枚举 */
package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
	)
	fmt.Printf("a = %d b = %d\n", a, b)

	const c = iota // 遇见const就会置0
	fmt.Printf("c = %d\n", c)

	const (
		d = iota // 后面可以不写iota
		e
		f
	)
	fmt.Printf("d = %d e = %d f = %d\n", d, e, f)

	const (
		j          = iota
		h1, h2, h3 = iota, iota, iota // 三个变量的值都一样
		i          = iota
	)
	fmt.Printf("j = %d h1 = %d h2 = %d h3 = %d i = %d\n", j, h1, h2, h3, i)

}

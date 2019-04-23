/* new函数 */
package main

import "fmt"

func main() {
	var p *int
	p = new(int) // 相当于c里的malloc，返回值是地址
	*p = 555
	fmt.Println("*p =", *p)

	q := new(int)
	*q = 666
	fmt.Println("*q =", *q)
}

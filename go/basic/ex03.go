/* 接受输入 */
package main

import "fmt"

func main() {
	var a int

	fmt.Println("enter a number:")
	fmt.Scanf("%d", &a) // 需要& 和c一样
	fmt.Printf("%d\n", a)
}

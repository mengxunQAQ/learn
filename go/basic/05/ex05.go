/* 字符串转换 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte, 0, 64)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1024, 10) // 指定10进制添加
	slice = strconv.AppendQuote(slice, "hello")

	fmt.Println(string(slice))
	fmt.Println(len(slice))

	// 整型转换成字符串
	str := strconv.Itoa(6666)
	fmt.Println(str)

	// 字符串转换整型
	a, _ := strconv.Atoi("567") // 如果不能转换，就会返回一个error，这里忽略掉了0_0
	fmt.Println(a)
}

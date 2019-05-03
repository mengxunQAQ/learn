/* 字符串的操作 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 验证是否包含字符串
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	// Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "x")
	fmt.Println("buf = ", buf)

	// Index 打印位置，不存在返回-1
	fmt.Println(strings.Index("abchello", "hello"))
	fmt.Println(strings.Index("abchello", "hi")) // -1

	// Repeat
	buf = strings.Repeat("go", 3)
	fmt.Println("buf =", buf) // gogogo

	// Split
	buf = "hello;abc;go"
	s2 := strings.Split(buf, ";")
	fmt.Println("s2 =", s2)

	// Trim
	buf = strings.Trim("    are u ok?   ", " ") // 去掉头尾的空格
	fmt.Printf("buf = #%s#\n", buf)

	// Fields 去掉空格把元素放入切片
	s3 := strings.Fields("   are u ok  ")
	fmt.Println("s3 =", s3)
}

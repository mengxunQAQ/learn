/* if判断 */
package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		fmt.Println(a)
	}

	if b := 1; b == 1 { // if支持带一个初始化语句，初始化语句和判断语句以分号分割
		fmt.Println(b)
	} else if b == 2 {
		fmt.Println(b)
	} else {
		fmt.Println(b)
	}

}

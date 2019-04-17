/*无参数有返回值的函数*/
package main

import "fmt"

func main() {
	a := myfunc01()
	fmt.Println(a)

	b := myfunc02()
	fmt.Println(b)

	c := myfunc03()
	fmt.Println(c)
}

func myfunc01() int {
	return 666
}

func myfunc02() (result int) {
	return 666
}

func myfunc03() (result int) {
	result = 666
	return
}

/* method
 方法：给某个类型绑定一个函数
 * type 重命名的类型不能是指针！
 * 如果两个方法名一样，但是绑定的类型不用，则视为两个不同的方法
 * 方法不支持重载
*******************************
type long *int
func (tmp long) test(){
}
错误
*******************************
type long int
func (tmp *long) test(){
}
正确
*/

package main

import "fmt"

type long int

// a是接受者，接受者也是参数
func (a long) Add02(b long) (result long) {
	result = a + b
	return
}

func main() {

	result1 := Add(1, 2)
	fmt.Println("result1 =", result1)

	var a long = 2
	result2 := a.Add02(3)
	fmt.Println("result2 =", result2)
}

func Add(a, b int) (result int) {
	result = a + b
	return
}

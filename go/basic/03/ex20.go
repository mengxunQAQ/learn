/* map的遍历 */
package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}

	// key -> 键, value -> 值
	for key, value := range m {
		fmt.Printf("%d =====> %s\n", key, value)
	}

	// value -> key所对应的值  ok -> key是否存在
	value, ok := m[1]
	if ok == true {
		fmt.Println("m[1] =", value)
	} else {
		fmt.Println("key不存在")
	}
}

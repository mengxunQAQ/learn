/* append扩容的特性 */
package main

import "fmt"

func main() {
	// append 当容量不足时，以 2 倍扩容
	s := []int{}
	fmt.Println("cap(s) =", cap(s))

	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println("cap(s) =", cap(s))

	s = append(s, 4)
	fmt.Println("cap(s) =", cap(s))
}

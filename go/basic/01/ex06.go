/* for循环 和 range */
package main

import "fmt"

func main() {
	str := "abc"

	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	for _, data := range str {
		fmt.Printf("%c\n", data)
	}

	for i := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	goto End
	fmt.Println("can not print")

End:
	fmt.Println("go to end")

}

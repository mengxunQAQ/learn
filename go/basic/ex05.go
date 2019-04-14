package main

import "fmt"

func main() {
	var num int

	fmt.Scanf("%d", &num)

	switch num {
	case 1:
		fmt.Println("enter 1...") // go默认有break
		fallthrough               // 相当于c不加break
	case 2:
		fmt.Println("enter 2...")
	case 3:
		fmt.Println("enter 3...")
	default:
		fmt.Println("out of range!")

	}
}

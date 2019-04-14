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
	case 3, 4: // 多个条件放一行
		fmt.Println("enter 3,4...")
	default:
		fmt.Println("out of range!")

	}

	a := 2

	switch { // 没有条件
	case a > 2: // 条件放在case后面
		fmt.Println("a > 2")
	case a > 1:
		fmt.Println("a > 1")
	default:
		fmt.Println("a <= 1")
	}

	switch b := 3; b { // 初始化
	case 3:
		fmt.Println("b = 3")
	case 2:
		fmt.Println("b = 2")
	default:
		fmt.Println("b = xxx")
	}
}

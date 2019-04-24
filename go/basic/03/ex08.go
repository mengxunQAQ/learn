/* 生成随机数 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需一次
	// 如果种子的参数一样，那么每次的随机数都一样
	rand.Seed(time.Now().UnixNano()) // 以当前时间做种子参数

	for i := 0; i < 5; i++ {
		//fmt.Println("rand =", rand.Int()) // 随机很大的数
		fmt.Println("rand =", rand.Intn(100)) // 限制在100内
	}
}

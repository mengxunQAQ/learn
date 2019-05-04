/* GOMAXPROCS */
package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(2)
	fmt.Println(n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}

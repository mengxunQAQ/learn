/* goroutine */
package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is new task")
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask()

	for {
		fmt.Println("this is main task")
		time.Sleep(time.Second)
	}
}

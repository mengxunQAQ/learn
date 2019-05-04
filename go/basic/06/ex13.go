/*
单向channel
* 正常的双向channel可以隐式转换为单向
* 但是，单向不能转换为双向
* a := make(chan int)
* var write chan<- int = a
*/
package main

func main() {
	var writeCh chan<- int
	//<-writeCh
	writeCh <- 1

	var readCh <-chan int
	<-readCh
}

/* error */
package main

import "fmt"
import "errors"

func main() {
	err1 := fmt.Errorf("%s", "this is err")
	fmt.Println(err1)

	err2 := errors.New("this is err")
	fmt.Println(err2)
}

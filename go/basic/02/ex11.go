/* 导包 */
package main // 必须

// 忽略此包
import _ "fmt"

func main() {

}

/*
// 给包起别名
import io "fmt"

func main() {
	io.Println("this is test")
}
*/

/*
// . 操作  容易和自定义的函数起命名冲突
import . "fmt" // 调用不需要重复写包名

//方式1
import "fmt" // 导入包，必须使用
import "os"

//方式2, 常用
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is test")
	fmt.Println("os.Args =", os.Args)
}
*/

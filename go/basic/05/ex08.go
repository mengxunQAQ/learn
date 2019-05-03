/* struct tag的使用 */
package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量必须【大写】
type Class struct {
	Name   string   `json:"name"` // 输出成name小写
	Member []string `json:"-"`    // 不输出
}

func main() {
	s := Class{Name: "3-1", Member: []string{"Alen", "Bob", "Clock"}}

	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Println("s =", string(buf))

}

/* json转普通类型 */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	tmp := make(map[string]interface{})
	jsonBuf := `
	{
		"name": "Alen",
		"age":  23
	}`

	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		panic(err)
	}

	fmt.Println("tmp =", tmp)
}

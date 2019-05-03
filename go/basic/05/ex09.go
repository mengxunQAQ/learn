/* map -> json */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 2)
	m["name"] = "Alen"
	m["age"] = 22

	result, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println("m =", string(result))
}

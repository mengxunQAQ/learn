/*类型判断*/
package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	i := make([]interface{}, 3)

	i[0] = 1
	i[1] = "abc"
	i[2] = Student{"mike", 22}

	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Println(index, "int:", value)
		} else if value, ok := data.(string); ok == true {
			fmt.Println(index, "string:", value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Println(index, "Student:", value)
		}
	}

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Println(index, value)
		case string:
			fmt.Println(index, value)
		case Student:
			fmt.Println(index, value)
		}
	}
}

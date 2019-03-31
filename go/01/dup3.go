package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[sting]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)


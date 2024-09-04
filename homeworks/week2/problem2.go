package main

import (
	"fmt"
)

var count map[string]int

func Problem2() {
	count = make(map[string]int)
	var text string
	fmt.Println("Enter the string you want: ")
	fmt.Scanln(&text)
	for _, v := range text {
		count[string(v)]++
	}
	fmt.Println("Result: ", count)
}

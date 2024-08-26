package main

import (
	"fmt"
)

func checkString(text string) bool {
	return len(text)%2 == 0
}

func Problem2() {
	var text string

	fmt.Println("Enter the string you want: ")
	fmt.Scanln(&text)

	fmt.Printf("Result: %t", checkString(text))
}

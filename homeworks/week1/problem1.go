package main

import (
	"fmt"
)

func perimeter(length, width float64) float64 {
	return 2 * (length + width)
}

func area(length, width float64) float64 {
	return length * width
}

func Problem1() {
	var length, width float64

	// Input length and width
	fmt.Println("Enter the length of rectangle")
	fmt.Scanln(&length)
	fmt.Println("Enter the width of rectangle")
	fmt.Scanln(&width)

	// Output perimeter and area
	fmt.Printf("Perimeter of rectangle %.2f\n", perimeter(length, width))
	fmt.Printf("Area of rectangle %.2f\n", area(length, width))
}

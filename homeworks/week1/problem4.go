package main

import (
	"fmt"
)

func twoSum(slice []int, target int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i]+slice[j] == target {
				fmt.Printf("Two numbers that sum up to target: %d, %d", slice[i], slice[j])
			}
		}
	}
}

func Problem4() {

	var length int
	fmt.Println("Enter the length of slice: ")
	fmt.Scanln(&length)
	fmt.Println("Enter elements of slice: ")

	// Input slice
	var slice []int
	slice = make([]int, 0)
	for i := 0; i < length; i++ {
		var element int
		fmt.Scanln(&element)
		slice = append(slice, element)
	}
	fmt.Println("Slice: ", slice)

	// Input target
	var target int
	fmt.Println("Enter the target: ")
	fmt.Scanln(&target)

	// Output two sum
	twoSum(slice, target)
}

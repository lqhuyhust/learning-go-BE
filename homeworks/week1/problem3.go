package main

import (
	"fmt"
	"sort"
)

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func Problem3() {
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

	// Output sorted slice
	sort.Ints(slice)
	fmt.Println("Sorted slice: ", slice)

	// Output sum
	fmt.Println("Sum of slice: ", sum(slice))

	// Output average
	fmt.Println("Average of slice: ", float64(sum(slice))/float64(length))

	// Output get max
	fmt.Println("Max of slice: ", slice[length-1])

	// Output get min
	fmt.Println("Min of slice: ", slice[0])

}

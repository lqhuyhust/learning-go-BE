package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	name string
	job  string
	year int
}

func Problem4() {
	f, err := os.Open("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	people := make([]Person, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// print raw data of file a.txt
		fmt.Println(scanner.Text())

		// split data to get detail data of each line
		parts := strings.Split(scanner.Text(), "|")
		name := strings.ToUpper(parts[0])
		job := strings.ToLower(parts[1])
		year, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		// add person data to slice
		person := Person{name, job, year}
		people = append(people, person)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(people)
}

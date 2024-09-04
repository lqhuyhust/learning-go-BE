package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	job  string
	year int
}

func (p *Person) getAge() int {
	return time.Now().Year() - p.year
}

func (p *Person) checkSuitable() bool {
	return p.year%len(p.name) == 0
}

func Problem1() {
	var name, job string
	var year int

	fmt.Println("Enter the name of the person")
	fmt.Scanln(&name)
	fmt.Println("Enter the job of the person")
	fmt.Scanln(&job)
	fmt.Println("Enter the born year of the person")
	fmt.Scanln(&year)

	person := Person{name, job, year}
	fmt.Println("The age of ", person.name, "is ", person.getAge())
	if person.checkSuitable() {
		fmt.Println("The person is suitable for work")
	} else {
		fmt.Println("The person is not suitable for work")
	}
}

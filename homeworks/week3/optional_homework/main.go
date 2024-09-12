package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	url        = "https://dummy.restapiexample.com/api/v1/employees"
	numWorkers = 5
)

type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeAge    int    `json:"employee_age"`
	EmployeeSalary int    `json:"employee_salary"`
	ProfileImage   string `json:"profile_image"`
}

type Response struct {
	Employees []Employee `json:"data"`
}

func getEmployees(url string) ([]Employee, error) {
	// Set dealay time and max retries in case error 429
	delay := 1 * time.Second
	maxRetries := 5

	for maxRetries > 0 {
		// Call API form url
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to get employees: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			// Read response body
			responseData, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to get response body: %w", err)
			}

			// Convert response body to struct
			var response Response
			err = json.Unmarshal(responseData, &response)
			if err != nil {
				return nil, fmt.Errorf("failed to convert employees: %w", err)
			}

			return response.Employees, nil
		} else if resp.StatusCode == http.StatusTooManyRequests {
			// Handle HTTP 429 by implementing exponential backoff
			fmt.Println("Too many requests, retrying in", delay, "seconds")
			time.Sleep(delay)
			delay *= 2 // Exponential backoff: double the delay time
			maxRetries--
			continue
		} else {
			return nil, fmt.Errorf("failed to get employees: status code %d", resp.StatusCode)
		}
	}

	return nil, fmt.Errorf("too many retries, giving up")

}

func worker(id int, jobs <-chan Employee, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started employee", j.EmployeeName)
		time.Sleep(time.Second)
		var result = j.EmployeeSalary / j.EmployeeAge
		fmt.Println("worker", id, "finished employee", j.EmployeeName, "with result", result)
		results <- result
	}
}

func createWorkerPool(employees []Employee) {
	// Create worker pool
	var numJobs = len(employees)
	jobs := make(chan Employee, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- employees[j-1]
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func main() {
	////////////////////////////////////
	// PROBLEM 1: Get slice of employees
	////////////////////////////////////
	fmt.Println("Problem 1: Get slice of employees")
	employees, err := getEmployees(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employees)
	}

	////////////////////////////////
	// PROBLEM 2: Create worker pool
	////////////////////////////////
	fmt.Println("Problem 2: Create worker pool")
	createWorkerPool(employees)
}

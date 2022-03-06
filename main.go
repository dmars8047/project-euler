package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dmars8047/eulertools"
	"github.com/dmars8047/project-euler/problem1"
	"github.com/dmars8047/project-euler/problem10"
	"github.com/dmars8047/project-euler/problem11"
	"github.com/dmars8047/project-euler/problem12"
	"github.com/dmars8047/project-euler/problem2"
	"github.com/dmars8047/project-euler/problem3"
	"github.com/dmars8047/project-euler/problem4"
	"github.com/dmars8047/project-euler/problem5"
	"github.com/dmars8047/project-euler/problem6"
	"github.com/dmars8047/project-euler/problem7"
	"github.com/dmars8047/project-euler/problem8"
	"github.com/dmars8047/project-euler/problem9"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("\n\nError: Please enter the problem number to be executed.\n\n")
		return
	}

	problemNumber, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("\n\nSupplied problem number: %s is not a number\n\n", os.Args[1])
		return
	}

	eulertools.PrintPremise(problemNumber)

	result := 0

	start := time.Now()

	switch problemNumber {
	case 1:
		result = problem1.Execute()
	case 2:
		result = problem2.Execute()
	case 3:
		result = problem3.Execute()
	case 4:
		result = problem4.Execute()
	case 5:
		result = problem5.Execute()
	case 6:
		result = problem6.Execute()
	case 7:
		result = problem7.Execute()
	case 8:
		result = problem8.Execute()
	case 9:
		result = problem9.Execute()
	case 10:
		result = problem10.Execute()
	case 11:
		result = problem11.Execute()
	case 12:
		result = problem12.Execute()
	default:
		fmt.Printf("\nError: The solution for Project Euler problem number: %d has not been added to theis program.\n\n", problemNumber)
		return
	}

	execTime := time.Since(start)

	fmt.Printf("\nResult: %d, Execution time: %f (Seconds)\n\n", result, execTime.Seconds())
}

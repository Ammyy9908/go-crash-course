package main

import "fmt"

func main() {
	fmt.Println("Starting Day 6: Conditional Statements in Go")

	// if
	// if else
	// if else if else
	// switch case

	// if
	if 5 > 2 {
		fmt.Println("5 is greater than 2")
	}

	// if else
	if 5 < 2 {
		fmt.Println("5 is less than 2")
	} else {
		fmt.Println("5 is greater than 2")
	}

	// if else if else
	if 5 < 2 {
		fmt.Println("5 is less than 2")
	} else if 5 > 2 {
		fmt.Println("5 is greater than 2")
	} else {
		fmt.Println("5 is equal to 2")
	}

	// switch case
	// switch case with default
	// switch case with fallthrough
	// switch case with multiple expressions
	// switch without a Condition

	// switch with Initialization

	// switch case
	condition := "A"
	switch condition {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
	}

	// switch case with default
	condition = "D"
	switch condition {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")

	default:
		fmt.Println("Default")
	}

	// switch case with fallthrough
	condition = "A"
	switch condition {
	case "A":
		fmt.Println("A")
		fallthrough
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")

	}

	// switch case with multiple expressions

	condition = "A"
	switch condition {
	case "A", "B":
		fmt.Println("A or B")
	case "C":
		fmt.Println("C")
	}

	// switch without a Condition
	condition = "A"

	switch {
	case condition == "A":
		fmt.Println("A")
	case condition == "B":

		fmt.Println("B")
	case condition == "C":
		fmt.Println("C")
	}

	// switch with Initialization
	switch condition := "A"; condition {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
	}
}

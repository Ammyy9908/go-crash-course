package main

import "fmt"

func main() {
	// 1. Arrays: fixed length list of things
	// 2. Slices: arrays with dynamic length
	// 3. Maps: key value pairs

	// Arrays
	var fruits = [2]string{"Apple", "Orange"}
	var numbers = [2]int{1, 2}
	var colors = [2]string{"Red", "Blue"}
	fmt.Print(fruits, numbers, colors)

	// Slices: slices with dynamic length
	var fruits_slice = []string{"Apple", "Orange"}
	var numbers_slice = []int{1, 2}
	var colors_slice = []string{"Red", "Blue"}
	fmt.Print(fruits_slice, numbers_slice, colors_slice)

	// Maps: key value pairs
	var fruits_map = map[string]string{"Apple": "Red", "Orange": "Orange"}
	var numbers_map = map[string]int{"One": 1, "Two": 2}
	var colors_map = map[string]string{"Red": "Red", "Blue": "Blue"}
	fmt.Print(fruits_map, numbers_map, colors_map)

	// Checking for Existence in map
	// 1. Check if the key exists
	// 2. Check if the value exists
	// 3. Check if the key value pair exists

	number, ok := numbers_map["One"]
	fmt.Println(number, ok)
	if !ok {
		fmt.Println("Key does not exist")
	} else {
		fmt.Println("Key exists")
	}

	//Iterating Over Maps
	for key, value := range numbers_map {
		fmt.Println(key, value)
	}

}

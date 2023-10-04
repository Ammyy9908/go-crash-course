package main

import "fmt"

func main() {
	fmt.Println("Welcome to Day 7: Loops and functions in Go")
	//loops
	//for loop
	//while loop
	//do while loop
	//infinite loop
	//nested loops
	//break
	//continue
	//functions
	//function with no parameters and no return type
	//function with parameters and no return type
	//function with parameters and return type
	//function with no parameters and return type
	//function with multiple return types
	//function with named return types
	//function with variadic parameters

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//do while loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	//nested loops
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}

	//break

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//continue

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//functions calls
	hello()
	hello_name("Sammy")
	fmt.Println(add(1, 2))
	fmt.Println(pi())
	fmt.Println(add_sub(1, 2))
	fmt.Println(add_sub_named(1, 2))
	fmt.Println(add_all(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

//function with no parameters and no return type

func hello() {
	fmt.Println("Hello")
}

//function with parameters and no return type

func hello_name(name string) {
	fmt.Println("Hello", name)
}

//function with parameters and return type

func add(a int, b int) int {
	return a + b
}

//function with no parameters and return type

func pi() float32 {
	return 3.14
}

//function with multiple return types

func add_sub(a int, b int) (int, int) {
	return a + b, a - b
}

//function with named return types

func add_sub_named(a int, b int) (add int, sub int) {
	add = a + b
	sub = a - b
	return
}

//function with variadic parameters

func add_all(a ...int) int {
	total := 0
	for _, num := range a {
		total += num
	}
	return total
}

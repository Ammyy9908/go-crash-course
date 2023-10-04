package main

import "fmt"

func main() {
	fmt.Println("Welcome to Day 4:Variables, Constants, and Basic Data Types in Go")
	//variables
	var name string = "John"

	//constants
	const pi float64 = 3.14

	//basic data types
	var age int = 20
	var isCool = true
	isCool = false

	//shorthand
	fname := "John"
	size := 1.3

	fmt.Println(name, age, isCool, fname, size, pi)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", fname)

	//all data types in golang
	//string
	user_name := "Sammy"
	fmt.Println(user_name)
	//bool
	is_valid := true
	fmt.Println(is_valid)
	//int
	age = 20
	fmt.Println(age)

	//int int8 int16 int32 int64

	//example for int,int8,int16,int32

	//int8 -128 to 127
	//int16 -32768 to 32767
	//int32 -2147483648 to 2147483647

	//uint8 0 to 255
	//uint16 0 to 65535
	//uint32 0 to 4294967295

	//uint uint8 uint16 uint32 uint64 uintptr

	//byte - alias for uint8
	//example
	var a uint8 = 255
	fmt.Println(a)

	//rune - alias for int32
	var Rune uint32 = 2147483647
	fmt.Println(Rune)

	//represents a Unicode code point

	//float32 float64
	var cgpa float32 = 9.5
	fmt.Println(cgpa)

	//complex64 complex128
	var c complex64 = 5 + 5i
	fmt.Println(c)

	//Type Interference: Go automatically assigns a type to a variable based on its value
	var b = true
	fmt.Println(b)

	//example
	fmt.Println(c)

	// Variables of Different Types: declaring types in a single block
	var (
		//integer
		age1 int = 20
		//float
		cgpa1 float32 = 9.5

		//boolean
		is_valid1 bool = true

		//string
		user_name1 string = "Sammy"
	)

	fmt.Println(age1, cgpa1, is_valid1, user_name1)
}

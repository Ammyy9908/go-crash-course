# Day 4: Variables, Constants, and Basic Data Types in Go

## Declaring and Initializing Variables

In Go, variables can be declared in multiple ways.

```go
 var name string
var age int = 25
city := "San Francisco"
```

The := operator is a shorthand for declaring and initializing a variable.

## Understanding Data Types

Go is a statically typed language, which means the type of a variable is checked at compile-time.
Some of the basic data types in Go are:
int and uint: Signed and unsigned integers of varying sizes (e.g., int8, uint8, int16, uint32, int64).
float32 and float64: Floating-point numbers.
bool: Boolean type with values true or false.
string: A sequence of characters.
complex64 and complex128: Complex numbers with float32 and float64 real and imaginary parts respectively.

## Using Constants

Constants are similar to variables, but their values cannot be changed after they are defined.
They are declared using the const keyword.

```go
const Pi = 3.14159
```

Constants can also be character, string, boolean, or numeric values.
Go supports a feature called iota for generating increasing numbers to create enumerated constants.

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
)
```

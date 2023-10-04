# Day 3: Basic Syntax in Go

## Go's Syntax Basics

Go's syntax is influenced by the C programming language but is designed to be more concise and safe.
Go uses braces { } to delimit blocks of code, similar to C, C++, and Java.
Semicolons are used to separate statements in many programming languages, but in Go, they are inserted automatically at the end of each line by the lexer.

## Understanding Go's main Function

The main function serves as the entry point for a Go application.

Every Go executable program must include a main function which resides in the main package. This is where the execution of the program begins.

The main function does not take any arguments and doesn't return any values.

```
package main

import "fmt"
func main() {
fmt.Println("Hello, World!")
}
```

## Writing and Running a Simple Go Program

To write a Go program, you typically use a .go extension for the source file.
You can run the program directly using the go run command followed by the filename. For the above example, you'd use go run hello.go to run the program and see the "Hello, World!" output.
Alternatively, you can compile the program into an executable using go build. This generates an executable file in the same directory. You can then run the executable directly.

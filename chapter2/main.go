// This is a package declaration. Every Go program must start with a package declaration. Packages are Go's way to organise and reuse code.
package main

// This is how we include code from other packages to use with our program; the `fmt` package (short for format) implements formatting for input and output.
import "fmt"

// This is a comment

// Functions are the building blocks of a Go program. They have inputs, outputs and a series of steps called statements which are executed in orger. All fuctions start with the keyword `func`, followed by the name of the function (`main` in this case), a list of zero or more "parameters" surrounded by parentheses, an optional return type and a "body" which is surround by curly braces. 
func main() {
	fmt.Println("Hello World")
}

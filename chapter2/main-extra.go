// This is a package declaration. Every Go program must start with a package declaration. Packages are Go's way to organise and reuse code.
package main

// This is how we include code from other packages to use with our program; the `fmt` package (short for format) implements formatting for input and output.
import "fmt"

// This is a comment

// Functions are the building blocks of a Go program. They have inputs, outputs and a series of steps called statements which are executed in orger. All fuctions start with the keyword `func`, followed by the name of the function (`main` in this case), a list of zero or more "parameters" surrounded by parentheses, an optional return type and a "body" which is surround by curly braces. 
// The `fmt.Println` statement is made of three components. FIrst we access another function inside the `fmt` package called `Println` (that's the `fmt.Println` bit, `Println` means Print Line). Then we create a new string that contains `Hello World` and invoke (also known as call or execute) that function with the string as the first and only argument.
func main() {
	fmt.Println("Hello World my name is CJ")
}

# Chapter 2
## Obligatory Hello World Program

### Introduction
Hello folks, welcome to the first lesson of this crashcourse. Today we're going to make the obligatory `Hello World` program, as every developer should do when trying a new programming language.

### Full Code
Here is the full code for a Hello World program in Go.
```go
package main

import "fmt"

// This is a comment

func main() {
	fmt.Println("Hello World")
}
```
### Walkthrough
Let's walk through what each bit means!

This is a package declaration. Every Go program must start with a package declaration. Packages are Go's way to organise and reuse code.
```go
package main
```

This is how we include code from other packages to use with our program; the `fmt` package (short for format) implements formatting for input and output.
```go
import "fmt"
```

This is how we comment in Go, we use a double slash, like this, `//`
```go
// This is a comment
```

Functions are the building blocks of a Go program. They have inputs, outputs and a series of steps called statements which are executed in orger. All fuctions start with the keyword `func`, followed by the name of the function (`main` in this case), a list of zero or more "parameters" surrounded by parentheses, an optional return type and a "body" which is surround by curly braces. 

The `fmt.Println` statement is made of three components. FIrst we access another function inside the `fmt` package called `Println` (that's the `fmt.Println` bit, `Println` means Print Line). Then we create a new string that contains `Hello World` and invoke (also known as call or execute) that function with the string as the first and only argument.

```go
func main() {
	fmt.Println("Hello World")
}
```


### Problems
- What is whitespace?

- What is a comment? What are the two ways of writing a comment?

- Our program began with package main. What would the files in the fmt package begin with?

- We used the Println function defined in the fmt package. If we wanted to use the Exit function from the os package what would we need to do?

- Modify the program we wrote so that instead of printing Hello World it prints Hello, my name is followed by your name.

// This file is in the package named "main"
package main

// Import fmt
import "fmt"

// Make it easier to understand the code by fixing the language's bad choice of type names.
type integer int

// This is a func named main
func main() {
	// Set x to integer(8)
	x := integer(8)

	// Set y to 7
	y := integer(7)

	// Call the magic func
	magic(&y)

	// Set x to y
	x = 7

	// Call fmt.Println on x and y
	fmt.Println(x, y)

	// It is very important to return so the computer doesn't keep running code after the end of your func
	return
}

// This is a func called magic
func magic(a *integer) {
	// Set *a to *a ^ 5
	*a = *a ^ 5

	// This line is very important!
	return
}

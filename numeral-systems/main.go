package main

import "fmt"

func main() {
	const name = "Juan"
	fmt.Printf("Hello, %s!\n", name)

	// Short declaration variable
	shortDeclaration := "Hello, World!"
	fmt.Println(shortDeclaration)

	// Variable declaration with type -> cero value
	var intCeroValue int //0
	var stringCeroValue string // ""
	fmt.Println(intCeroValue, stringCeroValue)

	// Print values in binary and hexadecimal
	age := 31

	fmt.Println("Age: ", age)
	fmt.Printf("Binary: %b\n", age)
	fmt.Printf("Hexadecimal: %#X\n", age)
}

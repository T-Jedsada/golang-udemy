package main

import "fmt"

func main() {
	a := 10
	b := "string"
	c := 13.99
	d := true

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	// dec
	fmt.Println(42)

	// hex
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	// binary
	fmt.Printf("%d - %b \n", 42, 42)
}

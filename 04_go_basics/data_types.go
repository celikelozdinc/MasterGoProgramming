package main

import "fmt"

func main() {
	// rune is the alias for int32
	var char rune = 'r'
	fmt.Printf("%T\n", char) // => int32

	// overflows uint8
	//var distance uint8
	//distance = 5 * 200
	//_ = distance

	// Type is pointer to int
	x := 1
	ptr := &x
	fmt.Printf("%T\n", ptr) // => *int

}

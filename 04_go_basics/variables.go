package main

import "fmt"

func main() {

	// Compile-time error
	//const x int = 5 //typed constant
	//const y float64 = 2.2 * x

	const x = 7 //untyped constant
	const y float64 = 3.1
	fmt.Println(x * y)

	//iota
	const (
		North = iota
		East
		South
		West
	)

	fmt.Printf("%v %v\n", East, West)

}

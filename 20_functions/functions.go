package main

import "fmt"

// sum returns ans anonymous functions which gets an int and returns an int
func sum(first int) func(second int) (result int) {
	return func(second int) (result int) {
		result = first + second
		return // => NAKED RETURN STATEMENT
	}
}

func foo() {
	fmt.Println("Hello from foo()!")
}

func baz() {
	fmt.Println("Hello from baz()!")
}

func variadicFunc(ints ...int) {
	ints[0] = 50
}

func main() {

	defer baz()

	// Go supports first-class functions
	function := sum(5)
	fmt.Println(function(7))
	fmt.Println(function(1))
	fmt.Println(function(3))

	// defer is executed in reverse order
	defer foo()

	//variadic functions
	mySlice := []int{0, 2, 4, 6}
	fmt.Printf("%+v\n", mySlice)
	variadicFunc(mySlice...) // => modifies underlying array
	fmt.Printf("%+v\n", mySlice)

}

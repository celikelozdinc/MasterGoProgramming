package main

import "fmt"

func printer(value int) func() {
	return func() {
		fmt.Printf("From anonymous func: Value is %d\n", value)
	}
}

func sum(ints ...int) int {
	sum := 0
	for _, value := range ints {
		sum += value
	}
	return sum
}

// uses naked return statment
func modifiedSum(ints ...int) (sum int) {
	sum = 0
	for _, value := range ints {
		sum += value
	}
	return
}

func main() {

	// Exercise 9
	p := printer(5)
	fmt.Printf("Type : %T\n", p)
	p()
	p()
	p()

	// Exercise 4
	fmt.Printf("Sum is %d\n", sum(5, 10, 15, 20))

	// Execise 5
	fmt.Printf("Sum is %d\n", modifiedSum(5, 10, 15, 20))

}

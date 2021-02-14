package main

import (
	"fmt"
)

func swap(first *float64, second *float64) {
	temp := *first
	*first = *second
	*second = temp

}

func main() {

	// Exercise 3
	x, y := 5.5, 8.8
	fmt.Printf("Before swapping : %.2f and %.2f\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping : %.2f and %.2f\n", x, y)

	// Exercise 1
	k := 10.10
	fmt.Printf("Adress of k : %p\n", &k)
	ptr := &k
	fmt.Printf("Type of pointer : %T, Value of pointer :%v\n", ptr, ptr)
	fmt.Printf("Address of pointer : %p, Value of variable :%v\n", &ptr, *ptr)

}

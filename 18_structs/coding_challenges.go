package main

import "fmt"

func main() {

	// Exercise 1
	type Person struct {
		name   string
		age    int
		colors []string
	}

	me := &Person{name: "Me", age: 30, colors: []string{"Blue", "Aqua"}}
	you := &Person{name: "You", age: 20, colors: []string{"White", "Black"}}
	fmt.Printf("%+v\n", *me)  //=> uses plus modifier
	fmt.Printf("%+v\n", *you) //=> uses plus modifier

	// Exercise 2
	Andrei := *you
	Andrei.colors = []string{"Red", "Yellow"}
	fmt.Printf("%+v\n", Andrei) //=> uses plus modifier
	fmt.Printf("%+v\n", *you)   //=> uses plus modifier

	// Exercise 3
	for ind, val := range me.colors {
		fmt.Printf("My %d.favourite color is %s\n", ind, val)
	}

}

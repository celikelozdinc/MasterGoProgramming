package main

import "fmt"

type second uint
type duration second

func main() {
	var d duration
	fmt.Printf("Type --> %T\n", d) // => duration

}

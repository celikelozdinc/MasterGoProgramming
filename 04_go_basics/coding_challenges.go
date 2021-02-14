package main

import (
	"fmt"
)

//version := 3.1
var version float64 = 3.1

type duration int // alias declaration

func main() {

	var a float64 = 7.1
	x, y := true, 3.7
	//a, x := 5.5, false
	a, z := 5.5, false
	_, _, _, _ = a, x, y, z

	//name := 'Golang'
	name := "Golang"
	fmt.Println(name)

	const i int = 7
	const b float64 = 7.1
	//const c = i * b
	const c = float64(i) * b

	//const noIPv6 = math.Pow(2, 128) // => compiles, runtime error
	// since function calls belongs to runtime while constants belogs to compile time

	// iota
	const (
		Jun = iota + 6
		July
		August
	)
	fmt.Printf("Jun is the %dth month of the year, while July is the %dth and August is the %dth\n", Jun, July, August)

	const t float64 = 1.422349587101
	fmt.Printf("With 4 precision --> %.4f\n", t)

	shape := "circle"
	r := 3.2
	const PI float64 = 3.14159
	circumference := 2 * PI * r

	fmt.Printf("Shape : %q\n", shape)
	fmt.Printf("Circle's circumference with radius %.2f is %.2f\n", r, circumference)

	var hour duration = 3600
	minute := 60
	//fmt.Println(hour != minute) // => does not even compile
	fmt.Println(hour != duration(minute)) // => does not even compile
}

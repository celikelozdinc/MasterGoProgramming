package main

import "fmt"

func main() {

	runeFirst, _ := 'a', 'b'
	fmt.Printf("Type :%T, Value: %#v, value %d\n", runeFirst, runeFirst, runeFirst) // => int32, 97, 97

	asciis := "Golang"
	fmt.Printf("Size of the variable full of ascii characters : %d\n", len(asciis)) // => 6 letters : 6 bytes

	asciiWithNoAsciis := "ÜzümÜzüm"
	fmt.Printf("Size of the variable with ascii and non-ascii characters : %d\n", len(asciiWithNoAsciis)) // => 12 bytes

	s1 := "Golang"
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

}

package main

import (
	"log"
	"strings"
)

func main() {

	// EXERCISE 6 : rune slice
	log.Println(strings.Repeat("@", 20))
	s := "üzüm"
	runes := []rune(s)
	log.Printf("Rune slice : %#v\n", runes)
	for i, v := range runes {
		log.Printf("Rune %c at index %d\n", v, i)
	}
	log.Println(strings.Repeat("@", 20))

	// EXERCISE 5 : byte slice
	log.Println(strings.Repeat("@", 20))
	byteSlice := []byte(s)
	log.Printf("Byte slice : %#v\n", byteSlice)
	for i, v := range byteSlice {
		log.Printf("Byte %b at index %d\n", v, i)
	}
	log.Println(strings.Repeat("@", 20))

	// EXERCISE 4
	log.Println(strings.Repeat("@", 20))
	s1 := "Go is cool!"
	x := s1[0]
	log.Println(x)
	//s1[0] = 'x'
	log.Println(len(s1)) // => #bytes in string
	log.Println(strings.Repeat("@", 20))

	// EXERCISE 2
	log.Println(strings.Repeat("@", 20))
	myRune := 'ü'
	myStr1, myStr2 := "z", "m"
	myResultingStr := string(myRune) + myStr1 + string(myRune) + myStr2 // need to convert from rune to string
	log.Printf("Resulting string containg runes : %s\n", myResultingStr)
	log.Println(strings.Repeat("@", 20))

}

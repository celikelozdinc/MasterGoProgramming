package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s1[:])  // => s1[0:len(s1)] : s1[0] is included, s1[len(s1)] did not
	fmt.Println(s1[2:]) // => s1[2:len(s1)] : s1[2] is included, s1[len(s1)] did not
	fmt.Println(s1[:3]) // => s1[0:3] : s1[0] is included, s1[3] did not

	s1 = append(s1, 200)
	fmt.Println(s1)

	s1 = append(s1, 400)
	fmt.Println(s1)

	newS := append(s1, 600)
	s1[0] = -600
	fmt.Println(s1)
	fmt.Println(newS)

	fmt.Println(strings.Repeat("#", 20))

	cars := []string{"Ford", "Honda", "Audi", "BMW", "Nissan"}

	cars1 := cars[0:3]
	fmt.Printf("cars1 : Len -> %d, Cap -> %d\n", len(cars1), cap(cars1))

	cars2 := cars[2:5]
	fmt.Printf("cars2 : Len -> %d, Cap -> %d\n", len(cars2), cap(cars2))

	fmt.Println(strings.Repeat("#", 20))
	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[:1], "X", "Y")
	fmt.Printf("Letters : %#v, len : %d, cap : %d\n", letters, len(letters), cap(letters))
	//fmt.Println(letters[3])
	//fmt.Println(letters[4])
	//fmt.Println(letters[5])
	fmt.Println(letters[3:4])
	fmt.Println(letters[3:5])
	fmt.Println(letters[3:6])

}

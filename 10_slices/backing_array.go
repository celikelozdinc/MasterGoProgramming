package main

import (
	"fmt"
	"strings"
)

func main() {

	//arr1 is the backing array of the slices
	arr1 := [...]int{10, 20, 30, 40, 50}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	//Modify backing array of slices
	arr1[1] = -80
	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println(strings.Repeat("#", 10))

	slice1[1] = -30
	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println(strings.Repeat("#", 20))

	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	//All of 3 slices shares same backing array
	s3[1] = -600
	fmt.Printf("s1 : %v\n", s1)
	fmt.Printf("s1 : Len -> %d, Cap -> %d\n", len(s1), cap(s1))
	fmt.Printf("s3 : %v\n", s3)
	fmt.Printf("s3 : Len -> %d, Cap -> %d\n", len(s3), cap(s3))
	fmt.Printf("s4 : %v\n", s4)
	fmt.Printf("s4 : Len -> %d, Cap -> %d\n", len(s4), cap(s4))

	fmt.Println(strings.Repeat("#", 20))

	cars := []string{"Ford", "Honda", "Audi"}
	newCars := []string{}
	newestCars := []string{}
	newCars = append(newCars, cars[0:2]...) // brand-new slice from existing slice
	newestCars = append(cars, "Fiat") // brand-new slice from existing slice
	cars[0] = "Nissan"
	cars[1] = "BMW"
	cars[2] = "Subaru"
	fmt.Printf("cars -> %v\n", cars)
	fmt.Printf("newcars -> %v\n", newCars)
	fmt.Printf("newestCars -> %v\n", newestCars)
}

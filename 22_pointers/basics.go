package main

import "fmt"

func changeMap(m map[string]int) {
	m["a"] = 100
	m["c"] = 200
}
func changeSlice(s []int) {
	s[0] = -1
	s[1] = -2
}

func main() {

	// Comparing pointers
	// Equal if they point to same variable
	x, y := 5, 5
	ptr1, ptr2 := &x, &y
	fmt.Println(*ptr1 == *ptr2) //=> true
	fmt.Println(ptr1 == ptr2)   //=> false
	ptr3 := &y
	fmt.Println(*ptr3 == *ptr2) //=> true
	fmt.Println(ptr3 == ptr2)   //=> true

	// For maps and slices, no need to pass pointers to functions
	// Slices and maps are not meant to be used with pointers
	myMap := map[string]int{"a": -1, "b": -2}
	fmt.Printf("My map before changing it : %+v\n", myMap)
	changeMap(myMap)
	fmt.Printf("My map after changing it : %+v\n", myMap)

	mySlice := []int{100, 200}
	fmt.Printf("My slice before changing it : %+v\n", mySlice)
	changeSlice(mySlice)
	fmt.Printf("My slice after changing it : %+v\n", mySlice)

}

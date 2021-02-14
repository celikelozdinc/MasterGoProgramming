package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Repeat("#", 20))
	n1 := []int{10, 20, 30, 40, 50}
	n2 := n1[0:3]
	n2[0] = 66 // they have same backing array
	fmt.Println(n1)
	fmt.Println(strings.Repeat("#", 20))

	fmt.Println(strings.Repeat("#", 20))
	s1 := []int{10, 20, 30, 40, 50}
	s2 := append(s1, 100)
	s2[0] = 66 // they don't have same backing array????
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(strings.Repeat("#", 20))

	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("Coding Exercise 8")
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(newYears, years[:3]...)            // first 3 elements
	newYears = append(newYears, years[len(years)-3:]...) // last 3 elements
	fmt.Println(newYears)
	fmt.Println(strings.Repeat("#", 20))

	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("Coding Exercise 7")
	friends := []string{"Marry", "John", "Diana", "Paul"}
	newFriends := []string{}
	newFriends = append(newFriends, friends...) // slices are not connected
	friends[0] = "FirstItem"
	newFriends[0] = "FirstElement"
	fmt.Println(friends)
	fmt.Println(newFriends)
	fmt.Println(strings.Repeat("#", 20))

	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("Coding Exercise 5")
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	var sum int = 0
	for index, value := range nums[1 : len(nums)-1] {
		fmt.Printf("slice[%d]=%d\n", index, value)
		sum += value
	}
	fmt.Printf("Sum -> %d\n", sum)
	fmt.Println(strings.Repeat("#", 20))

}

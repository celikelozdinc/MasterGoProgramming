package main

func main() {

	// ellipsis notation
	// appends all the elements of slice newNumbers to slice numbers
	numbers := []int{10, 20, 30}
	newNumbers := []int{80, 90}
	numbers = append(numbers, newNumbers...)



}

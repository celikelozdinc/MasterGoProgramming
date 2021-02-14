package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}
	myArray[2] = 6
	fmt.Println(myArray)

	a := 10
	//myArray[0] = a // => Compile-time error
	myArray[0] = float64(a)

	//myArray[3] = 10.10 // => out-of bounds

	nums := [...]int{30, -1, -6, 90, -6}

	for _, value := range nums {
		if value < 0 {
			fmt.Printf("Ignoring negative values, such as %d\n", value)
		} else {
			if value%2 == 0 {
				fmt.Printf("Positive even number, such as %d\n", value)
			} else {
				fmt.Printf("Positive odd number, such as %d\n", value)
			}
		}
	}

}

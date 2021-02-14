package main

import (
	"fmt"
	"time"
)

func main() {

	for number := 1; number <= 5000; number++ {
		if number%5 == 0 && number%7 == 0 {
			fmt.Printf("%d is divisible both by 5 and 7\n", number)
		}
	}

	for curYear := 1991; curYear <= time.Now().Year(); {
		fmt.Printf("Now, the year is : %d\n", curYear)
		curYear++
	}

	// switch statement
	age := 19

	switch true {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congrats!You have just became major!")
	default:
		fmt.Println("You are major!")
	}
}

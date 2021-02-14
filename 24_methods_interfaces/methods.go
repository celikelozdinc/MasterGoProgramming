package main

import (
	"fmt"
	"strings"
)

type names []string

func (n names) report() {
	for i, v := range n {
		fmt.Printf("[%d]=%s, ", i, v)
	}
	fmt.Println("\n\n")
}

type car struct {
	brand string
	price int
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func main() {
	friends := names{"Albus", "Severus", "Lupin"}
	// Both invokation are valid
	friends.report()
	names.report(friends)

	fmt.Println(strings.Repeat("@", 50))

	myCar := car{"Fiat", 80000}
	fmt.Printf("Before trying modification using changeCar1(): %+v\n", myCar)
	myCar.changeCar1("Audi", 700000)
	fmt.Printf("After trying modification changeCar1(): %+v\n", myCar)

	fmt.Println("\n\n")

	fmt.Printf("Before trying modification using changeCar2(): %+v\n", myCar)
	myCar.changeCar2("Audi", 700000)
	fmt.Printf("After trying modification changeCar2(): %+v\n", myCar)

}

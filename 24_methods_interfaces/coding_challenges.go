package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo, brand string
}

func (c car) License() string {
	fmt.Printf("Licence no of my vehicle : %q\n", c.licenceNo)
	return c.licenceNo
}

func (c car) Name() string {
	fmt.Printf("Brand of my vehicle : %q\n", c.brand)
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

type money float64

func (m money) print() {
	fmt.Printf("Money with 2 decimal points : %.2f\n", m)
}

func (m money) printStr() (result string) {
	result = fmt.Sprintf("%.2f", m)
	return //=> Naked return statement
}

type book struct {
	title string
	price float64
}

func (b book) changePrice(p float64) {
	b.price = p
}

func (b *book) modifyPrice(p float64) {
	b.price = p
}

func main() {

	// Exercise 1 (Interfaces)
	var veh vehicle
	newCar := car{"17 AAU 779", "Fiat"}
	veh = newCar
	fmt.Printf("Dynamic type of veh : %T\n", veh)
	fmt.Printf("Dynamic value of veh : %+v\n", veh)
	veh.License()
	veh.Name()

	//Exercise 3 (Interfaces)
	var x interface{}
	x = cube{edge: 5}
	//v := volume(x) //=> need a type assertion here
	v := volume(x.(cube))
	fmt.Printf("Cube volume :%v\n", v)

	// Exercise 1 (Methods)
	var myMoney money = 100.12345678
	myMoney.print()

	// Exercise 2 (Methods)
	fmt.Println(myMoney.printStr())

	// Exercise 4 (Methods)
	bestBook := book{title: "The Trial by Kafka", price: 9.9}
	bestBook.changePrice(11.99)
	fmt.Printf("Book's price : %.2f\n", bestBook.price) //=> no change
	bestBook.modifyPrice(11.99)
	fmt.Printf("Book's price : %.2f\n", bestBook.price) //=> updates

}

package main

import (
	"fmt"
	"log"
)

func main() {

	var myMap map[string]string
	if myMap == nil {
		log.Println("This is nil map")
	} else {
		log.Println("Not a nil map")
	}

	myMap1 := map[string]int{}
	if myMap1 == nil {
		log.Println("This is nil map")
	} else {
		log.Println("Not a nil map")
	}

	myMap1["FirstKey"] = 10
	myMap1["SecondKey"] = 20
	myMap1["FourthKey"] = 0
	log.Printf("Zero value for the key : %d\n", myMap1["Third"])

	//use comma okey idiom
	v, ok := myMap1["ThirdKey"]
	if ok {
		log.Printf("ThirdKey exist in the map : %d\n", v)
	} else {
		log.Printf("ThirdKey does not exist in the map")
	}

	v, ok = myMap1["FourthKey"]
	if ok {
		log.Printf("FourthKey exist in the map : %d\n", v)
	} else {
		log.Printf("FourthKey does not exist in the map")
	}

	// comparing by using Sprintf
	// if the both keys and values are string

	firstMap := map[string]string{"Key1": "Value1", "Key2": "Value2"}
	secondMap := map[string]string{"Key2": "Value2", "Key1": "Value1"}
	s1, s2 := fmt.Sprintf("%s", firstMap), fmt.Sprintf("%s", secondMap)

	if s1 == s2 {
		log.Println("Maps are equal")
	} else {
		log.Println("Maps are not equal")
	}

	// map header
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbours := friends
	friends["Dan"] = 200 // also updates neighbours, since their headers are same
	fmt.Printf("%#v ------- %#v\n",friends,neighbours)

	employees := map[string]int{} // empty and initialized map
	newEmp := employees // their headers are same
	newEmp["Andrei"] = 30
	fmt.Println(employees["Andrei"])

}

package main

import (
	"log"
	"strings"
)

func main() {

	// Exercise 3
	log.Println(strings.Repeat("@", 20))
	//m := map[int]bool{"1": true, 2: false, 3: false}
	m := map[int]bool{1: true, 2: false, 3: false}
	log.Printf("Map before deletion : %#v\n", m)
	delete(m, 2)
	log.Printf("Map after deletion : %#v\n", m)

	for i, v := range m {
		log.Printf("map[%d] == %v\n", i, v)
	}
	log.Println(strings.Repeat("@", 20))

	// Exercise 2
	log.Println(strings.Repeat("@", 20))
	var m1 map[int]bool // nil map, uninitialized
	//m1[5] = true // => panic : assignment to entry in nil map
	_ = m1

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}
	//log.Println(m2 == m3) // => map can only be compared to nil
	_, _ = m2, m3
	log.Println(strings.Repeat("@", 20))

	// Exercise 1
	log.Println(strings.Repeat("@", 20))
	var map1 map[int]int // nil map
	log.Printf("Type : %T, Value : %#v\n", map1, map1)

	map2 := map[int]string{} // empty map
	map2[10] = "Abba"

	//by using comam okey idiom
	value, ok := map2[10]
	if ok {
		log.Printf("Key exists, its value %s\n", value)
	} else {
		log.Println("Key does not exist")
	}

	value, ok = map2[100]
	if ok {
		log.Printf("Key exists, its value %s\n", value)
	} else {
		log.Println("Key does not exist")
	}
	log.Println(strings.Repeat("@", 20))

}

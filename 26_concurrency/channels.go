package main

import "fmt"

func main() {

	// uninitialized unbuffered channel
	var ch chan int
	fmt.Println(ch) //=> its zero value its nil
	ch = make(chan int)
	fmt.Println(ch) // => stores address
	defer close(ch)

	// declare + initialize unbuffered BIDIRECTIONAL CHANNEL
	chann := make(chan int)
	defer close(chann)

	// only for receiving (unbuffered, unidirectional)
	chan1 := make(<-chan string)
	//defer close(chan1)

	// only for sending (unbuffered, unidirectional)
	chan2 := make(chan<- int)
	//defer close(chan2)

	buffered := make(chan int, 3) //=> buffered channel

	fmt.Printf("%T, %T, %T, %T\n", chann, chan1, chan2, buffered)

}

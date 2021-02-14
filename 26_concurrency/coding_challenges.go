package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(first float64, second float64, wg *sync.WaitGroup) {
	sum := first + second
	fmt.Printf("Summation : %.2f\n", sum)
	wg.Done()
}

func power(num int, c chan int) {
	result := num * num
	c <- result
}

func main() {

	// Exercise 1
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("Mr, Wick", &wg)
	wg.Wait()

	// Exercise 2
	var waitG sync.WaitGroup
	waitG.Add(3)
	go sum(1.111, 2.222, &waitG)
	go sum(4.444, 5.555, &waitG)
	go sum(7.777, 8.888, &waitG)
	waitG.Wait()

	// Exercise 3 => Anonymous functions
	n := 6.4
	var wGroup sync.WaitGroup
	wGroup.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f = %.2f\n", f, x)
		wg.Done()
	}(n, &wGroup)
	wGroup.Wait()

	// Exercise 4
	var waitGroup sync.WaitGroup
	waitGroup.Add(50)
	for number := 100; number < 150; number++ {
		go func(f int, wg *sync.WaitGroup) {
			x := math.Sqrt(float64(f))
			fmt.Printf("Square root of %d = %v\n", f, x)
			wg.Done()
		}(number, &waitGroup)
	}
	waitGroup.Wait()

	// Exercise 1
	var c1 chan float64      //=> bidirectional unbuffered channel
	c2 := make(<-chan rune)  // => receive-only channel
	c3 := make(chan<- rune)  // => send-only channel
	c4 := make(chan int, 10) // => bidirectional buffered channel with capacity of 10 ints
	fmt.Printf("%T,%T,%T,%T\n", c1, c2, c3, c4)

	// Exercise 2
	channel := make(chan string)
	go func(s string, c chan string) {
		c <- s // send
	}("Hello From main goroutine", channel)
	fmt.Println(<-channel) // receive

	// Exercise 3
	//c := make(<-chan int) // receive-only
	c := make(chan int)
	go func(n int) {
		c <- n // send
	}(100)
	fmt.Println(<-c) // receive

	// Exercise 4

	ch := make(chan int)
	for num := 1; num < 51; num++ {
		go power(num, ch) // send
		fmt.Println(<-ch) // receive
	}

	// Exercise 5 => using anonymous function
	for num := 100; num < 1003; num++ {
		go func(num int, c chan int) {
			result := num * num
			c <- result // send
		}(num, ch)
		fmt.Println(<-ch) // receive
	}

}

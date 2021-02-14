package main

import "fmt"

func init() {
	fmt.Println("Hello from init func #1")
}

func init() {
	fmt.Println("Hello from init func #2")
}

func main() {
	fmt.Println("Hello from main func")
}

func init() {
	fmt.Println("Hello from init func #3")
}

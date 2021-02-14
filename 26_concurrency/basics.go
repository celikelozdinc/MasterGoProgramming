package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("O/S :%s\n", runtime.GOOS)
	fmt.Printf("Arch :%s\n", runtime.GOARCH)
	fmt.Printf("# CPUs : %d\n", runtime.NumCPU())             // => logical CPUs usable by the current process
	fmt.Printf("# Goroutines : %d\n", runtime.NumGoroutine()) //=> number of goroutines that currently exist => in this case, only main goroutine
	fmt.Printf("GOMAXPROCS : %d\n", runtime.GOMAXPROCS(0))
}

package main

import (
	"fmt"
	"time"
)

<<<<<<< HEAD
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
}

=======
>>>>>>> cb95587042fe112a9834fcc903d6e22211b5d421
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
<<<<<<< HEAD
=======

func main() {
	go spinner(100 * time.Millisecond)
	const n = 50
	fibN := fib(n)
	fmt.Printf("\rFibonachi(%d) = %d\n", n, fibN)
}
>>>>>>> cb95587042fe112a9834fcc903d6e22211b5d421

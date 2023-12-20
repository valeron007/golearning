package fib

import (
	"os"
	"strconv"
)

func fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	println(fib(n))
}

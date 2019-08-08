package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// fmt.Println(factorial(5))

	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
}

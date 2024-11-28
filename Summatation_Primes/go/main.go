package main

import "fmt"

func isPrime(a int) bool {
	if a < 2 {
		return false
	}

	if a == 2 {
		return true
	}

	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var sum int
	for i := 1; i < 2000000; i += 1 {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Printf("Sumation of primes under two million: %d\n", sum)
}

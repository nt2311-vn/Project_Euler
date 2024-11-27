package main

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	if x == 2 {
		return true
	}

	i := 2

	for i*i <= x {
		if x%i == 0 {
			return false
		}

		i += 1
	}

	return true
}

func main() {
	n, i := 0, 2

	for {
		if isPrime(i) {
			n++

			if n == 10001 {
				fmt.Println("Result:", i)
				break
			}
		}

		i++
	}
}

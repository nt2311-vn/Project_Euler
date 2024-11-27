package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	root := math.Sqrt(float64(x))

	for i := 2; i < int(root); i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	const target = 600_851_475_143
	sqrt_target := math.Sqrt(float64(target))

	for i := int(sqrt_target); i >= 2; i-- {
		if isPrime(i) && target%i == 0 {
			fmt.Println(i)
			break
		}
	}
}

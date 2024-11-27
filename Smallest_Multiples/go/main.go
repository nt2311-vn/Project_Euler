package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	root := math.Sqrt(float64(x))

	if x <= 2 {
		return true
	}

	for i := 2; i <= int(root); i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func findPrimeFactor(x int) []int {
	primeFactor := make([]int, 0)

	if x <= 2 {
		return primeFactor
	}

	for i := 2; i <= x; i++ {
		if isPrime(i) && x%i == 0 {
			primeFactor = append(primeFactor, i)
		}
	}

	return primeFactor
}

func toPrimeFactor(x int) map[int]int {
	rs := make(map[int]int)
	num := x

	primeFactor := findPrimeFactor(num)

out:
	for i := 0; i < len(primeFactor); i++ {
		for num != 1 {
			if num%primeFactor[i] == 0 {
				rs[primeFactor[i]]++
				num /= primeFactor[i]
			} else {
				continue out
			}
		}
	}

	return rs
}

func main() {
	mapPrime := make(map[int]int)
	nonPrime := make([]int, 0)

	for i := 2; i <= 20; i++ {
		if isPrime(i) {
			mapPrime[i]++
		} else {
			nonPrime = append(nonPrime, i)
		}
	}

	for _, num := range nonPrime {
		mapNum := toPrimeFactor(num)

		for k, v := range mapNum {
			if mapPrime[k] < v {
				mapPrime[k] = v
			}
		}
	}
	smallest := 1

	for k, v := range mapPrime {
		smallest *= int(math.Pow(float64(k), float64(v)))
	}

	fmt.Println(smallest)
}

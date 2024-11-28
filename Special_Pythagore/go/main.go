package main

import "fmt"

func isSpecialTriplet(a, b, c int) bool {
	return a > 0 && a < b && b < c && a*a+b*b == c*c && a+b+c == 1000
}

func main() {
	for c := 333; c < 450; c++ {
		for b := 216; b < c && b < 386; b++ {
			a := 1000 - b - c

			if a <= 0 {
				break
			}

			if isSpecialTriplet(a, b, c) {
				fmt.Printf("Found special triplet a: %d, b: %d, c: %d\n", a, b, c)
				fmt.Printf("Result: %d\n", a*b*c)
				break
			}
		}
	}
}

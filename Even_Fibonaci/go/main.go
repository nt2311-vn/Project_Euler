package main

import "fmt"

func main() {
	a, b := 1, 1
	c, sum := 0, 0

	for b < 4_000_000 {
		c = a + b
		a, b = b, c

		if c%2 == 0 {
			sum += c
		}
	}

	fmt.Println(sum)
}

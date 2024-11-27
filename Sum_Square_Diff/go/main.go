package main

import "fmt"

func sumOfSquare(x int) int {
	var sum int
	for i := 1; i <= x; i++ {
		sum += i * i
	}

	return sum
}

func squareOfSum(x int) int {
	var sum int

	for i := 1; i <= x; i++ {
		sum += i
	}

	return sum * sum
}

func main() {
	diff := squareOfSum(100) - sumOfSquare(100)

	fmt.Println(diff)
}

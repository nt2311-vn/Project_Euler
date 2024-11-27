package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	numStr := strconv.Itoa(x)

	for i, j := 0, len(numStr)-1; i < j; i, j = i+1, j-1 {
		if numStr[i] != numStr[j] {
			return false
		}
	}

	return true
}

func main() {
	palindromes := make([]int, 0)
	for i := 999; i >= 100; i-- {
		var product int

		for j := 999; j >= 100; j-- {
			product = i * j

			if isPalindrome(product) {
				palindromes = append(palindromes, product)
			}
		}
	}

	maxNum := 0

	for _, num := range palindromes {
		if num > maxNum {
			maxNum = num
		}
	}

	fmt.Println("Result:", maxNum)
}

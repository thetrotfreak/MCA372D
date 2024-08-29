package main

import (
	"fmt"
	"math"
	"strings"
)

// Task 1: Find pairs of integers in the list whose sum is equal to the target
func findPairs(numbers []int, target int) [][]int {
	seen := make(map[int]bool) // track seen numbers
	var result [][]int

	for _, num := range numbers {
		complement := target - num

		// Check if the complement has already been seen, but also consider its reflection
		if seen[complement] || seen[num] {
			continue
		}

		seen[num] = true
		result = append(result, []int{num, complement})
	}

	return result
}

// Task 2: Generate Fibonacci series up to a given number recursively
func generateFibonacciSeries(limit int) []int {
	fibonacci := []int{0, 1}

	for i := 2; ; i++ {
		nextFib := fibonacci[i-1] + fibonacci[i-2]
		if nextFib > limit {
			break
		}
		fibonacci = append(fibonacci, nextFib)
	}

	return fibonacci
}

// Task 3: Check whether a given string is a palindrome or not
func isPalindrome(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, ",", "")

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

// Task 4: Generate prime numbers up to a given limit
func generatePrimes(limit int) []int {
	primes := []int{}

	for i := 2; i <= limit; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	// Task 1
	numbers := []int{2, 7, 11, 15, 3, 6, 8, 12}
	target := 18
	pairs := findPairs(numbers, target)
	fmt.Printf("Pairs with sum %d are: %v\n", target, pairs)

	// Task 2
	fibonacciLimit := 50
	fibonacciSeries := generateFibonacciSeries(fibonacciLimit)
	fmt.Printf("Fibonacci series up to %d: %v\n", fibonacciLimit, fibonacciSeries)

	// Task 3
	palindromeStr := "race,car"
	fmt.Printf("Is \"%s\" a palindrome? %v\n", palindromeStr, isPalindrome(palindromeStr))

	// Task 4
	primeLimit := 20
	primes := generatePrimes(primeLimit)
	fmt.Printf("Prime numbers up to %d: %v\n", primeLimit, primes)
}

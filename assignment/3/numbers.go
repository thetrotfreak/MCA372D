package main

import (
	"fmt"
	"strings"
)

// FindPairs function takes a slice of integers and a target sum, returning pairs that sum to the target
func findPairs(numbers []int, target int) [][]int {
	seen := make(map[int]bool) // track seen numbers
	var result [][]int

	for _, num := range numbers {
		complement := target - num

		// Check if the complement has already been seen, considering the smaller and larger values
		if seen[complement] && complement > num {
			continue
		}

		seen[num] = true
		result = append(result, []int{num, complement})
	}

	return result
}

// fibonacci function returns the nth Fibonacci number recursively
func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("invalid fibonacci number: %d", n)
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}
	a, err := fibonacci(n - 1)
	if err != nil {
		return 0, err
	}
	b, err := fibonacci(n - 2)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}

// isPalindrome checks if a string is a palindrome, ignoring special characters and case
func isPalindrome(s string) bool {
	s = removeSpecialChars(s) // remove special characters
	s = strings.ToLower(s)    // convert to lowercase

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// removeSpecialChars removes non-alphanumeric characters from a string
func removeSpecialChars(s string) string {
	var result string
	for _, char := range s {
		if isAlphaNumeric(char) {
			result += string(char)
		}
	}
	return result
}

// isAlphaNumeric checks if a character is alphanumeric (a-z, A-Z, 0-9)
func isAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

// isPrime checks if a number is prime using efficient square root check
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// generatePrimes returns a slice of prime numbers up to a given limit
func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	// Replace with user input
	numbers := []int{2, 7, 11, 15, 3, 6, 8, 12}
	target := 18

	pairs := findPairs(numbers, target)
	fmt.Println("Pairs with sum", target, "are:", pairs)

	// Get fibonacci number (replace with user input)
	n := 8
	fib, err := fibonacci(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The", n, "th Fibonacci number is:", fib)
	}

	// Check palindrome (replace with user input)
	s := "A man, a plan, a canal: Panama"
	if isPalindrome(s) {
		fmt.Println(s, "is a palindrome")
	} else {
		fmt.Println(s, "is not a palindrome")
	}

	// Generate prime numbers
	fmt.Println("Primes till 100 are", generatePrimes(100))
}

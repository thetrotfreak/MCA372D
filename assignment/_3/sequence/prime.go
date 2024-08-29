package sequence

import (
	"math"
)

func Primes(limit int) []int {
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
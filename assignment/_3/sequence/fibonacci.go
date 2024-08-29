package sequence

func Fibonacci(limit int) []int {
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
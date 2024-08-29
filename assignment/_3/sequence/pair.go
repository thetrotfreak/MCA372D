package sequence

func FindAll(numbers []int, target int) [][]int {
	pairs := [][]int{}
	visited := make(map[int]bool)

	for _, num := range numbers {
		complement := target - num
		if visited[complement] && !visited[num] {
			pairs = append(pairs, []int{complement, num})
			visited[complement] = true
			visited[num] = true
		} else {
			visited[num] = false
		}
	}

	return pairs
}
package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	payload := `[
		{"name": "Rob", "age": 68},
		{"name": "Ken", "age": 81},
		{"name": "Robert", "age": 59},
		{"name": "Matt", "age": 44},
		{"name": "Jon", "age": 53}
	]`

	var users []User

	if err := json.Unmarshal([]byte(payload), &users); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	var wg sync.WaitGroup
	// number of users = number of goroutines
	// each goroutine will add an User's Age
	wg.Add(len(users))

	var sum int
	var mutex sync.Mutex

	for _, user := range users {
		go func(u User) {
			defer wg.Done()

			// lock sum
			// to prevent race
			mutex.Lock()
			sum += u.Age
			mutex.Unlock()
		}(user)
		// schedule the goroutine right away
	}

	wg.Wait()
	averageAge := float64(sum) / float64(len(users))
	fmt.Printf("Average age of users: %.2f\n", averageAge)
}

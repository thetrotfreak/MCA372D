package main

import (
	"bufio"
	"fmt"
	"os"
)

func consolidateGrades(m map[string]float64) (float64, string) {
	var sum float64
	var grade string

	for _, score := range m {
		sum += score
	}
	avg := sum / float64(len(m))

	switch {
	case avg >= 90:
		grade = "A"
	case avg >= 80:
		grade = "B"
	case avg >= 70:
		grade = "C"
	case avg >= 60:
		grade = "D"
	default:
		grade = "F"
	}
	return avg, grade
}

func readGrades(m map[string]map[string]float64) {
	scanner := bufio.NewScanner(os.Stdin)

	var total_students int
	fmt.Print("Total students? ")
	fmt.Scanf("%d", &total_students)
	for i := 1; i <= total_students; i++ {
		fmt.Print("Student? ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("something went wrong!")
		}
		student := scanner.Text()

		var total_subjects int
		fmt.Print("Total subjects? ")
		fmt.Scanf("%d", &total_subjects)

		m[student] = make(map[string]float64)
		for j := 1; j <= total_subjects; j++ {
			fmt.Print("Subject? ")
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				fmt.Println("something went wrong!")
			}
			subject := scanner.Text()

			var score float64
			fmt.Print("Score? ")
			fmt.Scanf("%f", &score)

			m[student][subject] = score
		}
		average, grade := consolidateGrades(m[student])
		fmt.Printf("Gradings for \"%v\":\n\tAverage : %.2f\n\tGrade   : %v\n", student, average, grade)
	}
}

func main() {
	// assumption
	// student is a key to a value of subject:score k:v pair
	var grades = make(map[string]map[string]float64)
	readGrades(grades)
}
